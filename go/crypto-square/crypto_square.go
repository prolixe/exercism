package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

const testVersion = 2

func normalize(input string) (output string) {
	for _, r := range input {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			output += string(r)
		}

	}
	output = strings.ToLower(output)
	return
}

func getRectangle(length int) (row, col int) {

	// Get the approximate square side
	row = int(math.Sqrt(float64(length)))
	col = row
	for col*row < length {
		if col == row {
			col++
		} else if col > row {
			row++
		}
	}
	return
}

func Encode(input string) (output string) {
	normalized := normalize(input)
	r, c := getRectangle(len(normalized))
	for j := 0; j < c; j++ {
		for i := 0; i < r; i++ {
			if j+i*c >= len(normalized) {
				break
			}
			output += string(normalized[j+i*c])
		}
		output += " "
	}
	output = strings.Trim(output, " ")
	return
}
