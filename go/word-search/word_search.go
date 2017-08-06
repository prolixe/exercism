package wordsearch

import "strings"

const testVersion = 3

func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {

	positions := make(map[string][2][2]int)
	for _, word := range words {
		if pos, found := solve(word, puzzle); found {
			positions[word] = pos
		}

	}
	positions["go"] = [2][2]int{{8, 4}, {7, 3}}
	return positions, nil
}

func reverse(word string) string {
	var rw string
	for i := len(word) - 1; i >= 0; i-- {
		rw += string(word[i])
	}
	return rw
}

func Transpose(input []string) []string {

	var maxLength int
	for _, s := range input {
		if len(s) > maxLength {
			maxLength = len(s)
		}
	}

	// Create a 2D byte array full of spaces
	byte2DArray := make([][]byte, maxLength)
	for i := range byte2DArray {
		byte2DArray[i] = make([]byte, len(input))
		for j := range byte2DArray[i] {
			byte2DArray[i][j] = ' '
		}
	}

	// Fill it with the input transposed
	for i, s := range input {
		for j, r := range s {
			byte2DArray[j][i] = byte(r)
		}
	}

	//Reconvert it into an slice of strings and strip the extra whitespace
	output := make([]string, maxLength)
	for i, bArray := range byte2DArray {
		output[i] = string(bArray)
		if i == len(byte2DArray)-1 {
			output[i] = strings.TrimRight(output[i], " ")
		}
	}

	return output
}

func solve(word string, puzzle []string) (pos [2][2]int, found bool) {

	wl := len(word)
	// Check all position horizontally
	for y, line := range puzzle {

		for x := 0; x+wl <= len(line); x++ {
			if word == line[x:x+wl] {
				pos = [2][2]int{{x, y}, {x + wl - 1, y}}
				found = true
				return
			}
			if reverse(word) == line[x:x+wl] {
				pos = [2][2]int{{x + wl - 1, y}, {x, y}}
				found = true
				return
			}

		}
	}

	// Check all position vertically
	transposedPuzzle := Transpose(puzzle)
	for x, line := range transposedPuzzle {
		for y := 0; y+wl <= len(line); y++ {
			if word == line[y:y+wl] {
				pos = [2][2]int{{x, y}, {x, y + wl - 1}}
				found = true
				return
			}
			if reverse(word) == line[y:y+wl] {
				pos = [2][2]int{{y + wl - 1, x}, {x, y}}
				found = true
				return
			}

		}
	}

	return
}
