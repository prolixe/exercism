package igpay

import "strings"

const testVersion = 1

func PigLatin(input string) string {
	words := strings.Split(input, " ")

	pigWords := make([]string, 0)
	for _, word := range words {
		var consonnant string
	Loop:
		for i := 0; i < len(word); i++ {
			switch word[i] {
			case 'a', 'e', 'i', 'o', 'u':
				consonnant = word[:i]
				// special case that ignore the vowel if the precedent letter is a 'q'
				if len(consonnant) == 0 || consonnant[len(consonnant)-1] != 'q' {
					break Loop // Breaking with a label is nice.
				}

			}
		}
		switch {
		case len(consonnant) == 0:
			// Start with a vowel
			pigWords = append(pigWords, word+"ay")
		case len(consonnant) == 1:
			// Normal pig latin
			pigWords = append(pigWords, word[1:]+consonnant+"ay")
		case len(consonnant) > 1:
			// special cases
			if consonnant[0] == 'x' || consonnant[0] == 'y' {
				pigWords = append(pigWords, word+"ay")
			} else {
				pigWords = append(pigWords, word[len(consonnant):]+consonnant+"ay")
			}
		}
	}
	return strings.Join(pigWords, " ")
}
