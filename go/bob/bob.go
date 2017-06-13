package bob

import "strings"

const testVersion = 3

func Hey(s string) string {

	s = strings.Trim(s, " ")
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	switch {
	case strings.ToUpper(s) == s && strings.ContainsAny(strings.ToLower(s), alphabet):
		return "Whoa, chill out!"
	case len(s) > 0 && strings.EqualFold(string(s[len(s)-1:]), "?"):
		return "Sure."
	case !strings.ContainsAny(s, "abcdefghijklmnopqrstuvwxyz1234567890"):
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}
}
