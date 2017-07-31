package binarysearch

import "fmt"

const testVersion = 1

// Search is copied from the standard library
func Search(n int, f func(int) bool) int {
	i, j := 0, n
	for i < j {
		h := i + (j-i)/2 // i/2 + j/2
		if !f(h) {
			i = h + 1
		} else {
			j = h
		}
	}
	return i

}

// SearchInts is copied from the standard library
func SearchInts(a []int, x int) int {
	return Search(len(a), func(i int) bool { return a[i] >= x })
}

// Message indicate if and where the key was found
func Message(a []int, x int) string {
	i := SearchInts(a, x)
	if len(a) == 0 {
		return "slice has no values"
	}
	switch {
	case i == 0 && a[i] == x:
		return fmt.Sprintf("%d found at beginning of slice", x)
	case i == len(a)-1 && x == a[i]:
		return fmt.Sprintf("%d found at end of slice", x)
	case 0 < i && i < len(a)-1 && a[i] == x:
		return fmt.Sprintf("%d found at index %d", x, i)
	case 0 < i && i < len(a) && x > a[i-1] && x < a[i]:
		return fmt.Sprintf("%d > %d at index %d, < %d at index %d", x, a[i-1], i-1, a[i], i)
	case i == 0 && x < a[i]:
		return fmt.Sprintf("%d < all values", x)
	case i == len(a) && x > a[i-1]:
		return fmt.Sprintf("%d > all %d values", x, len(a))
	}
	return ""
}
