package strain

const testVersion = 1

type Ints []int
type Strings []string
type Lists [][]int

func (i Ints) Keep(f func(int) bool) (kept Ints) {
	for _, e := range i {
		if f(e) {
			kept = append(kept, e)
		}
	}
	return

}
func (i Ints) Discard(f func(int) bool) Ints {
	return i.Keep(func(i int) bool { return !f(i) })
}

func (s Strings) Keep(f func(string) bool) (kept Strings) {
	for _, e := range s {
		if f(e) {
			kept = append(kept, e)
		}
	}
	return
}

func (s Strings) Discard(f func(string) bool) Strings {
	return s.Keep(func(s string) bool { return !f(s) })
}

func (l Lists) Keep(f func([]int) bool) (kept Lists) {
	for _, e := range l {
		if f(e) {
			kept = append(kept, e)
		}
	}
	return

}
func (l Lists) Discard(f func([]int) bool) Lists {
	return l.Keep(func(l []int) bool { return !f(l) })
}
