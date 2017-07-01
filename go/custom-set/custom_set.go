package stringset

import "fmt"

const testVersion = 4

type Set map[string]bool

func New() Set {
	return make(Set)
}

func (set Set) String() string {
	s := "{"
	for key := range set {
		s += fmt.Sprintf("\"%s\", ", key)
	}
	if len(s) > 1 {
		s = s[:len(s)-2]
	}
	s += "}"
	return s
}

func NewFromSlice(stringSlice []string) Set {
	set := make(Set)
	for _, s := range stringSlice {
		set[s] = true
	}
	return set
}

func (set Set) IsEmpty() bool {
	return len(set) == 0
}
func (set Set) Has(s string) bool {
	return set[s]
}

func Subset(s1, s2 Set) bool {
	for e1 := range s1 {
		if !s2[e1] {
			return false
		}
	}
	return true
}

func Disjoint(s1, s2 Set) bool {
	for e1 := range s1 {
		if s2[e1] {
			return false
		}
	}
	return true
}

func Equal(s1, s2 Set) bool {
	if len(s1) != len(s2) {
		return false
	}

	for e2 := range s2 {
		if !s1[e2] {
			return false
		}
	}

	for e1 := range s1 {
		if !s2[e1] {
			return false
		}

	}
	return true
}

func (set Set) Add(s string) {
	set[s] = true
}

func Intersection(s1, s2 Set) Set {
	set := make(Set)
	for e1 := range s1 {
		if s2[e1] {
			set[e1] = true
		}
	}
	return set
}

func Difference(s1, s2 Set) Set {
	set := make(Set)
	for e1 := range s1 {
		set[e1] = true
	}

	for e2 := range s2 {
		if s1[e2] {
			delete(set, e2)
		}
	}
	return set
}

func Union(s1, s2 Set) Set {
	set := make(Set)
	for e1 := range s1 {
		set[e1] = true
	}

	for e2 := range s2 {
		set[e2] = true
	}
	return set
}
