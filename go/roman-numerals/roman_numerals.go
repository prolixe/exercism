package romannumerals

import "errors"

const testVersion = 3

type RomanArabic struct {
	roman  string
	arabic int
}

var baseNumerals = [...]RomanArabic{
	{"M", 1000},
	{"CM", 900},
	{"D", 500},
	{"CD", 400},
	{"C", 100},
	{"XC", 90},
	{"L", 50},
	{"XL", 40},
	{"X", 10},
	{"IX", 9},
	{"V", 5},
	{"IV", 4},
	{"I", 1},
}

func ToRomanNumeral(n int) (romanNumeral string, err error) {
	if n < 1 || n > 3999 {
		return "", errors.New("Invalid input")
	}

	for _, bn := range baseNumerals {
		for n >= bn.arabic {
			romanNumeral += bn.roman
			n -= bn.arabic
		}
	}
	return
}
