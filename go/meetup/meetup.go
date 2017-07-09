package meetup

import (
	"time"
)

const testVersion = 3

type WeekSchedule int

const (
	First WeekSchedule = iota
	Second
	Third
	Fourth
	Last
	Teenth
)

func Day(ws WeekSchedule, weekday time.Weekday, month time.Month, year int) int {
	date := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	switch ws {
	case First:
		for date.Weekday() != weekday {
			date = date.AddDate(0, 0, 1)
		}
		return date.Day()
	case Second:
		return Day(First, weekday, month, year) + 7
	case Third:
		return Day(First, weekday, month, year) + 14
	case Fourth:
		return Day(First, weekday, month, year) + 21
	case Last:
		date = date.AddDate(0, 1, -1) // Go the the last day of the month
		for date.Weekday() != weekday {
			date = date.AddDate(0, 0, -1)
		}
		return date.Day()

	case Teenth:
		for date.Weekday() != weekday || !(12 < date.Day() && date.Day() < 20) {
			date = date.AddDate(0, 0, 1)
		}
		return date.Day()
	}
	return 0
}