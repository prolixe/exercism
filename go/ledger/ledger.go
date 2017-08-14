package ledger

import (
	"bytes"
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const testVersion = 4

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

type EntryChan struct {
	i int
	s string
	e error
}

type Locale string

const (
	nlNL Locale = "nl-NL"
	enUS Locale = "en-US"
)

type Currency int

const (
	USD Currency = iota
	EUR
)

func NewCurrency(currency string) (Currency, error) {
	switch currency {
	case "USD":
		return USD, nil
	case "EUR":
		return EUR, nil
	default:
		return Currency(-1), errors.New("invalid currency")
	}
}

func SortEntries(e []Entry) func(i, j int) bool {
	return func(i, j int) bool {
		ei, ej := e[i], e[j]
		if ei.Date == ej.Date {
			// If the date is the same, compare the description
			if ei.Description == ej.Description {
				// if the description is the same, compare by change
				return ei.Change < ej.Change
			}
			return ei.Description < ei.Description
		}
		return ei.Date < ej.Date
	}
}

func createHeader(l Locale) (string, error) {
	var s string
	switch l {
	case nlNL:
		s = fmt.Sprintf("%-10s | %-25s | %s\n", "Datum", "Omschrijving", "Verandering")
	case enUS:
		s = fmt.Sprintf("%-10s | %-25s | %s\n", "Date", "Description", "Change")
	default:
		return "", errors.New("Invalid locale")
	}
	return s, nil
}

var dateRegex = regexp.MustCompile(`(\d{4})-(\d{1,2})-(\d{1,2})`)

func formatDate(date string, locale Locale) (string, error) {
	if !dateRegex.MatchString(date) {
		return "", errors.New("invalid date format")
	}
	results := dateRegex.FindStringSubmatch(date)
	year, month, day := results[1], results[2], results[3]

	var formattedDate string
	switch locale {
	case nlNL:
		formattedDate = fmt.Sprintf("%s-%s-%s", day, month, year)
	case enUS:
		formattedDate = fmt.Sprintf("%s/%s/%s", month, day, year)
	default:
		return "", errors.New("invalid locale")
	}
	return formattedDate, nil
}

func formatChange(change int, locale Locale, currency string) (string, error) {
	var b bytes.Buffer

	var currencySymbol string
	currencyType, err := NewCurrency(currency)
	if err != nil {
		return "", err
	}
	switch currencyType {
	case USD:
		currencySymbol = "$"
	case EUR:
		currencySymbol = "€"
	default:
		return "", errors.New("invalid currency")
	}
	// format the change
	negative := false
	if change < 0 {
		negative = true
		change *= -1
	}

	var sepCent, sepCurrency string
	var negativeFormat string = " "
	switch locale {
	case nlNL:
		if negative {
			negativeFormat = "-"
		}
		sepCent = ","
		sepCurrency = "."
	case enUS:
		if negative {
			b.WriteString("(")
			negativeFormat = ")"
		}
		sepCent = "."
		sepCurrency = ","
	default:
		return "", errors.New("invalid locale")
	}

	cents := change % 100

	b.WriteString(currencySymbol)
	if locale == nlNL {
		b.WriteString(" ")
	}
	// Insert a separation after each 3 digit
	rest := strconv.Itoa(change / 100)
	for i, r := range rest {
		b.WriteRune(r)
		if i%3 == 0 && len(rest) > 3 && i != len(rest)-1 {
			b.WriteString(sepCurrency)
		}
	}
	b.WriteString(sepCent)
	b.WriteString(fmt.Sprintf("%02d", cents))
	b.WriteString(negativeFormat)

	return b.String(), nil

}

func formatDescription(de string) string {
	if len(de) > 25 {
		de = de[:22] + "..."
	} else {
		de = de + strings.Repeat(" ", 25-len(de))
	}

	return de
}

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	var entriesCopy []Entry
	for _, e := range entries {
		entriesCopy = append(entriesCopy, e)
	}
	if len(entries) == 0 {
		if _, err := FormatLedger(currency, "en-US", []Entry{{Date: "2014-01-01", Description: "", Change: 0}}); err != nil {
			return "", err
		}
	}

	sort.Slice(entriesCopy, SortEntries(entriesCopy))

	s, err := createHeader(Locale(locale))
	if err != nil {
		return "", err
	}
	// Parallelism, always a great idea
	co := make(chan EntryChan)
	for i, et := range entriesCopy {
		go func(i int, entry Entry) {
			description := formatDescription(entry.Description)

			date, err := formatDate(entry.Date, Locale(locale))
			if err != nil {
				co <- EntryChan{e: err}
				return
			}
			change, err := formatChange(entry.Change, Locale(locale), currency)
			if err != nil {
				co <- EntryChan{e: err}
			}
			s := fmt.Sprintf("%-10s | %-25s | %13s\n", date, description, change)
			co <- EntryChan{i: i, s: s}
		}(i, et)
	}
	ss := make([]string, len(entriesCopy))
	for range entriesCopy {
		v := <-co
		if v.e != nil {
			return "", v.e
		}
		ss[v.i] = v.s
	}
	s += strings.Join(ss, "")
	return s, nil
}
