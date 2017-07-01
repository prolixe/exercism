package prime

const testVersion = 2

func Sieve(max int64) []int64 {
	primes := make([]bool, max)
	for i := int64(0); i < max; i++ {
		primes[i] = true
	}
	primes[0] = false
	primes[1] = false
	primes[2] = true

	var i int64
	for {
		for i < max && !primes[i] {
			i++
		}
		for j := int64(i << 1); j < max; j += i {
			primes[j] = false
		}
		i++
		if i > max {
			break
		}
	}

	primesList := make([]int64, 0)
	for i, p := range primes {
		if p {
			primesList = append(primesList, int64(i))
		}
	}

	return primesList
}

var primes = Sieve(894120)

func Factors(i int64) []int64 {
	factors := make([]int64, 0)
	count := 0 // to avoid looping forever, reliable up to 2^100
	for i > 1 || count > 100 {
		for _, p := range primes {
			if i%p == 0 {
				factors = append(factors, p)
				i /= p
				break
			}
		}
		count++
	}
	return factors
}
