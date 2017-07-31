package cipher

import (
	"regexp"
	"strings"
	"unicode"
)

const testVersion = 1

type Caeser int
type Vigenere string

var validVigenere = regexp.MustCompile(`^[a-z]+$`)
var invalidVigenere = regexp.MustCompile(`^a+$`)

func NewCaesar() Cipher {
	return Caeser(3)
}

func NewShift(shift int) Cipher {
	if shift == 0 || shift < -25 || shift > 25 {
		return nil
	}
	return Caeser(shift)
}

func NewVigenere(s string) Cipher {
	if !validVigenere.MatchString(s) || invalidVigenere.MatchString(s) {
		return nil
	}
	return Vigenere(s)
}

func (c Caeser) Encode(s string) string {
	return stringShift(s, int(c))
}

func (c Caeser) Decode(s string) string {
	return stringShift(s, -int(c))
}
func (v Vigenere) Encode(s string) string {
	return v.shift(s, true)
}

func (v Vigenere) Decode(s string) string {
	return v.shift(s, false)
}

func (v Vigenere) shift(s string, encode bool) (output string) {
	normalized := []byte(normalize(s))
	for i := range normalized {
		var c string
		if encode {
			c = stringShift(string(normalized[i]), int(byte(v[len(output)%len(v)])-'a'))
		} else {
			c = stringShift(string(normalized[i]), -int(byte(v[len(output)%len(v)])-'a'))
		}
		if len(c) == 1 {
			output += c
		}
	}
	return
}

func stringShift(s string, shift int) string {
	normalized := []byte(normalize(s))
	for i := range normalized {
		if normalized[i]+byte(shift) < 'a' {
			normalized[i] += byte(shift) + byte(26)
		} else if normalized[i]+byte(shift) > 'z' {
			normalized[i] += byte(shift) - byte(26)
		} else {
			normalized[i] += byte(shift)
		}
	}
	return string(normalized)
}

func normalize(input string) (output string) {
	for _, r := range input {
		if unicode.IsLetter(r) {
			output += strings.ToLower(string(r))
		}
	}
	return
}
