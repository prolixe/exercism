package wordy

import (
	"regexp"
	"strconv"
	"strings"
)

const testVersion = 1

type op int

const (
	invalid op = iota
	plus
	minus
	times
	divided
)

var regexNum = regexp.MustCompile("-?[0-9]+")
var regexOp = regexp.MustCompile("[^0-9\\-\\?]+")

func Answer(question string) (int, bool) {
	matchesNum := regexNum.FindAllString(question, -1)
	matchesOp := regexOp.FindAllString(question, -1)
	matchesOp = matchesOp[1:]

	if len(matchesNum) < 2 {
		return 0, false
	}
	a, _ := strconv.Atoi(matchesNum[0])
	for i := range matchesOp {
		b, _ := strconv.Atoi(matchesNum[i+1])
		switch getOp(matchesOp[i]) {
		case plus:
			a += b
		case minus:
			a -= b
		case times:
			a *= b
		case divided:
			a /= b
		}

	}

	return a, true
}

func getOp(s string) op {

	switch {
	case strings.Contains(s, "plus"):
		return plus
	case strings.Contains(s, "minus"):
		return minus
	case strings.Contains(s, "multiplied"):
		return times
	case strings.Contains(s, "divided"):
		return divided
	default:
		return invalid
	}
}
