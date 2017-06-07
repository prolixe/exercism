package leap

const testVersion = 3

// IsLeapYear returns a boolean value given an int representing a year
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
