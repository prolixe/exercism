package acronym

import "strings"

const testVersion = 3

func Abbreviate(s string) string {

	var a string
	words := strings.Split(strings.Replace(s, "-", " ", -1), " ")
	for _, word := range words {
		a += strings.ToUpper(string(word[0]))
	}
	return a

}
