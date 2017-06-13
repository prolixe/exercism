package pangram

import "strings"

const testVersion = 1

func IsPangram(s string) bool {

	sLower := strings.ToLower(s)
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	for _, c := range alphabet {
		if !strings.ContainsAny(sLower, string(c)) {
			return false
		}
	}
	return true

}
