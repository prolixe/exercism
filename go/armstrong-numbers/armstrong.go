package armstrong

import (
	"math"
	"strconv"
)

func IsNumber(n int) bool {

	initial := n
	powerSum := 0

	lenNumber := len(strconv.Itoa(n))

	for n > 0 {
		powerSum += int(math.Pow(float64(n%10), float64(lenNumber)))
		n /= 10
	}
	return powerSum == initial
}
