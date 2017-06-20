package isogram

import (
	"strings"
	"unicode"
)

const testVersion = 1

func IsIsogram(s string) bool {
	letterCount := make(map[rune]int)

	for _, l := range strings.ToLower(s) {
		if !unicode.IsLetter(l) {
			continue
		}
		letterCount[l]++
		if letterCount[l] > 1 {
			return false
		}
	}
	return true
}
