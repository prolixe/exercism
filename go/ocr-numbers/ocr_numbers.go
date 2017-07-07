package ocr

import (
	"fmt"
	"strings"
)

const testVersion = 1

func recognizeDigit(input string) string {

	switch input {
	case ` _ 
| |
|_|
   `:
		return "0"
	case `   
  |
  |
   `:
		return "1"
	case ` _ 
 _|
|_ 
   `:
		return "2"
	case ` _ 
 _|
 _|
   `:
		return "3"
	case `   
|_|
  |
   `:
		return "4"
	case ` _ 
|_ 
 _|
   `:
		return "5"
	case ` _ 
|_ 
|_|
   `:
		return "6"
	case ` _ 
  |
  |
   `:
		return "7"
	case ` _ 
|_|
|_|
   `:
		return "8"
	case ` _ 
|_|
 _|
   `:
		return "9"
	default:
		return "?"
	}
}

func Recognize(input string) []string {

	digits := make([]string, 0)
	lines := strings.Split(input, "\n")
	lines = lines[1:] // The first line in the test case is irregular, remove it.
	for i := 0; i < len(lines)/4; i++ {
		var line string
		for j := 0; j < len(lines[0]); j += 3 {
			d := []string{lines[i*4][j : j+3], lines[1+i*4][j : j+3], lines[2+i*4][j : j+3], lines[3+i*4][j : j+3]}
			fmt.Println(d)
			ocrDigit := strings.Join(d, "\n")
			line += recognizeDigit(ocrDigit)
			fmt.Println(line)
		}
		digits = append(digits, line)
	}
	return digits
}
