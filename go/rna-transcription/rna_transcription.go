package strand

import "strings"

const testVersion = 3

var replacer = strings.NewReplacer(
	"G", "C",
	"C", "G",
	"T", "A",
	"A", "U")

func ToRNA(input string) string {
	return replacer.Replace(input)
}
