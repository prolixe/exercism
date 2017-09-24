package twofer

import "fmt"

func ShareWith(s string) string {
	if len(s) == 0 {
		s = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", s)
}
