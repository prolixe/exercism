package beer

import (
	"errors"
	"fmt"
	"strings"
)

const testVersion = 1

func Verse(verseNum int) (string, error) {
	if verseNum < 0 || verseNum > 99 {
		return "", errors.New("Invalid verse")
	}

	if verseNum == 2 {

		return "2 bottles of beer on the wall, 2 bottles of beer.\nTake one down and pass it around, 1 bottle of beer on the wall.\n", nil
	}
	if verseNum == 1 {
		return "1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n", nil
	}
	if verseNum == 0 {
		return "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n", nil
	}

	verse := fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottles of beer on the wall.\n",
		verseNum,
		verseNum,
		verseNum-1)

	return verse, nil
}

func Verses(upperVerse, lowerVerse int) (string, error) {
	verses := make([]string, 0)
	for v := upperVerse; v >= lowerVerse; v-- {
		verse, err := Verse(v)
		if err != nil {
			return "", err
		}
		verses = append(verses, verse)
	}
	if len(verses) == 0 {
		return "", errors.New("Invalid input")
	}

	return strings.Join(verses, "\n") + "\n", nil
}

func Song() string {
	v, _ := Verses(99, 0)
	return v
}
