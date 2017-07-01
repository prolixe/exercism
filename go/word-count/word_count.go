package wordcount

import (
	"regexp"
	"strings"
)

const testVersion = 3

type Frequency map[string]int

var wordRegex = regexp.MustCompile("[[:alnum:]]+'?[[:alnum:]]+|[[:alnum:]]+")

func WordCount(phrase string) Frequency {

	p := strings.ToLower(phrase)
	words := wordRegex.FindAllString(p, -1)
	freq := make(Frequency)
	for _, w := range words {
		freq[w]++
	}

	return freq
}
