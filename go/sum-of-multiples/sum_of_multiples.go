package summultiples

const testVersion = 2

func SumMultiples(limit int, divisors ...int) (sum int) {
	for i := 0; i < limit; i++ {
		for _, d := range divisors {
			if i%d == 0 {
				sum += i
				break
			}
		}
	}
	return sum
}
