package pythagorean

const testVersion = 1

type Triplet [3]int

func (t Triplet) Perimeter() int {
	return t[0] + t[1] + t[2]
}

func Range(min, max int) (triplets []Triplet) {
	for a := min; a <= max; a++ {
		for b := a; b <= max; b++ {
			for c := b; c <= max; c++ {
				if a*a+b*b == c*c {
					triplets = append(triplets, Triplet{a, b, c})
				}
			}
		}
	}
	return
}

func Sum(p int) (perimeterTriplets []Triplet) {

	for _, t := range Range(1, p) {
		if t.Perimeter() == p {
			perimeterTriplets = append(perimeterTriplets, t)
		}
	}
	return
}
