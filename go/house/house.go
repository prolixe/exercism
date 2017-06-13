package house

import "strings"

const testVersion = 1

var verse = []string{

	"the house that Jack built.",
	`the malt
that lay in`,
	`the rat
that ate`,
	`the cat
that killed`,
	`the dog
that worried`,
	`the cow with the crumpled horn
that tossed`,
	`the maiden all forlorn
that milked`,
	`the man all tattered and torn
that kissed`,
	`the priest all shaven and shorn
that married`,
	`the rooster that crowed in the morn
that woke`,
	`the farmer sowing his corn
that kept`,
	`the horse and the hound and the horn
that belonged to`,
}

func Song() string {

	verses := make([]string, 0)
	for i := range verse {
		verses = append(verses, Verse(i+1))
	}
	return strings.Join(verses, "\n\n")

}

func recursiveVerse(v int) string {
	if v > 1 {
		return verse[v-1] + " " + recursiveVerse(v-1)
	}
	if v == 1 {
		return verse[v-1]
	}
	return ""
}

func Verse(v int) string {
	return "This is " + recursiveVerse(v)
}
