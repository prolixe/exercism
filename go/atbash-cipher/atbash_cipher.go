package atbash

import "strings"

const testVersion = 2

var replacer = strings.NewReplacer(" ", "", ".", "", ",", "")

func Atbash(input string) string {
	inputProcessed := replacer.Replace(input)
	inputProcessed = strings.ToLower(inputProcessed)
	output := make([]byte, 0)
	for i, r := range inputProcessed {
		if i%5 == 0 && i != 0 {
			output = append(output, ' ')
		}
		if byte('a') <= byte(r) && byte(r) <= byte('z') {
			output = append(output, byte('z'-(byte(r)-'a')))
		} else {
			output = append(output, byte(r))
		}
	}
	return string(output)
}
