package anagram

import "strings"

const testVersion = 2

type letterCount map[rune]int

func Detect(subject string, candidates []string) (anagrams []string) {

	subjectLower := strings.ToLower(subject)
	sc := countLetters(subject)
	for _, c := range candidates {
		if len(subjectLower) != len(c) {
			continue
		}

		cc := countLetters(c)
		if strings.ToLower(c) != subjectLower && isSame(sc, cc) {
			anagrams = append(anagrams, c)
		}
	}
	return
}

func countLetters(s string) letterCount {
	lc := make(letterCount)
	for _, r := range strings.ToLower(s) {
		lc[r]++
	}
	return lc
}

func isSame(a, b letterCount) bool {
	if len(a) != len(b) {
		return false
	}
	for key, value := range a {
		if b[key] != value {
			return false
		}
	}
	for key, value := range b {
		if a[key] != value {
			return false
		}
	}
	return true
}
