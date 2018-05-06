package sublist

import "reflect"

type Relation string

const (
	equal     Relation = "equal"
	sublist   Relation = "sublist"
	superlist Relation = "superlist"
	unequal   Relation = "unequal"
)

func Sublist(a, b []int) Relation {

	if len(a) == len(b) {
		e := true
		for i := 0; len(a) > i; i++ {
			if a[i] != b[i] {
				e = false
				break
			}
		}
		if e {

			return equal
		}
	}

	if len(a) == 0 {
		return sublist
	}

	if len(b) == 0 {
		return superlist
	}

	for i := 0; len(a) > i; i++ {
		for j := i; len(a) >= j; j++ {
			if reflect.DeepEqual(a[i:j], b) {
				return superlist
			}
		}
	}

	for i := 0; len(b) > i; i++ {
		for j := i; len(b) >= j; j++ {
			if reflect.DeepEqual(b[i:j], a) {
				return sublist
			}
		}
	}

	return unequal
}
