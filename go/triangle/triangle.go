package triangle

import "math"

const testVersion = 3

type Kind int

const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

func isSideValid(s float64) bool {
	return s != 0 && !math.IsNaN(s) && !math.IsInf(s, 0)
}

// Organize your code for readability.
func KindFromSides(a, b, c float64) Kind {

	if !isSideValid(a) || !isSideValid(b) || !isSideValid(c) {
		return NaT
	}

	// Find the largest side
	hypotenuse := math.Max(math.Max(a, b), c)
	if a+b+c < 2*hypotenuse {
		// the 2 shortest side can't reach each other, impossible triangle
		return NaT
	}

	if a == b && b == c {
		return Equ
	}
	if a == b || a == c || b == c {
		return Iso
	}

	return Sca

}
