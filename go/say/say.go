package say

import (
	"math"
	"strings"
)

const testVersion = 1

var m = map[uint64]string{0: "",
	1: "one", 2: "two", 3: "three", 4: "four", 5: "five", 6: "six", 7: "seven", 8: "eight",
	9: "nine", 10: "ten", 11: "eleven", 12: "twelve", 13: "thirteen",
	14: "fourteen", 15: "fifteen", 16: "sixteen", 17: "seventeen",
	18: "eighteen", 19: "nineteen", 20: "twenty",
	30:   "thirty",
	40:   "forty",
	50:   "fifty",
	60:   "sixty",
	70:   "seventy",
	80:   "eighty",
	90:   "ninety",
	100:  "hundred",
	1000: "thousand",
	1e6:  "million",
	1e9:  "billion",
	1e12: "trillion",
	1e15: "quadrillion",
	1e18: "quintillion",
}

func Say(n uint64) string {
	if n == 0 {
		return "zero"
	}
	return strings.Trim(recursiveSay(n), " ")
}

func recursiveSay(n uint64) string {

	switch {
	case n < 21:
		return m[n]
	case n < 100:
		return m[n/10*10] + "-" + recursiveSay(n%10)
	case n < 1000:
		return m[n/100] + " " + m[100] + " " + recursiveSay(n%100)
	case n < 1e6:
		return recursiveSay(n/1e3) + " " + m[1e3] + " " + recursiveSay(n%1e3)
	case n < 1e9:
		return recursiveSay(n/1e6) + " " + m[1e6] + " " + recursiveSay(n%1e6)
	case n < 1e12:
		return recursiveSay(n/1e9) + " " + m[1e9] + " " + recursiveSay(n%1e9)
	case n < 1e15:
		return recursiveSay(n/1e12) + " " + m[1e12] + " " + recursiveSay(n%1e12)
	case n < 1e18:
		return recursiveSay(n/1e15) + " " + m[1e15] + " " + recursiveSay(n%1e15)
	case n <= math.MaxUint64:
		return recursiveSay(n/1e18) + " " + m[1e18] + " " + recursiveSay(n%1e18)
	}
	return ""
}
