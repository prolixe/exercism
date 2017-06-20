package foodchain

import (
	"fmt"
	"strings"
)

const testVersion = 3

var lyricsComments = map[string]string{
	"fly":    "",
	"spider": "It wriggled and jiggled and tickled inside her.\n",
	"bird":   "How absurd to swallow a bird!\n",
	"cat":    "Imagine that, to swallow a cat!\n",
	"dog":    "What a hog, to swallow a dog!\n",
	"goat":   "Just opened her throat and swallowed a goat!\n",
	"cow":    "I don't know how she swallowed a cow!\n",
	"horse":  "She's dead, of course!",
}

var animal = [...]string{"fly", "spider", "bird", "cat", "dog", "goat", "cow", "horse"}

func Verse(v int) string {

	verse := fmt.Sprintf("I know an old lady who swallowed a %s.\n", animal[v-1])
	verse += lyricsComments[animal[v-1]]
	if v == 8 {
		return verse
	}
	// Why section
	for i := v; i > 1; i-- {
		verse += fmt.Sprintf("She swallowed the %s to catch the %s.\n", animal[i-1], animal[i-2])
		if i == 3 && v != 2 {
			// Spider is special
			verse = verse[:len(verse)-2]
			verse += " that wriggled and jiggled and tickled inside her.\n"
		}
	}

	verse += "I don't know why she swallowed the fly. Perhaps she'll die."
	return verse

}

func Verses(start, end int) string {
	verses := make([]string, end-start+1)
	for i := start; i <= end; i++ {
		verses[i-1] = Verse(i)
	}
	return strings.Join(verses, "\n\n")
}

func Song() string {
	return Verses(1, 8)
}
