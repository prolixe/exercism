package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

const testVersion = 2

var double = map[byte]byte{
	'1': '2',
	'2': '4',
	'3': '6',
	'4': '8',
	'5': '1',
	'6': '3',
	'7': '5',
	'8': '7',
	'9': '9',
	'0': '0',
}

func Valid(input string) bool {

	doubledInput := []byte(strings.Replace(input, " ", "", -1))
	if len(doubledInput) < 2 {
		return false
	}
	for _, b := range doubledInput {
		if !unicode.IsDigit(rune(b)) {
			return false
		}
	}

	for i := len(doubledInput) - 2; i >= 0; i -= 2 {
		doubledInput[i] = double[doubledInput[i]]
	}

	return sumOfDigits(doubledInput)%10 == 0
}

func sumOfDigits(bs []byte) (sum int) {
	for _, b := range bs {
		if d, ok := strconv.Atoi(string(b)); ok == nil {
			sum += d
		}
	}
	return
}
