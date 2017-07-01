package transpose

import "strings"

const testVersion = 1

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
