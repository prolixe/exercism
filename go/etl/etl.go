package etl

import "strings"

const testVersion = 1

func Transform(input map[int][]string) map[string]int {

	output := make(map[string]int)
	for key, slice := range input {
		for _, letter := range slice {
			output[strings.ToLower(string(letter))] = key
		}
	}
	return output

}
