package prime

const testVersion = 1

func Sieve(max int) []int {
	primes := make([]bool, max)
	for i := 0; i < max; i++ {
		primes[i] = true
	}
	primes[0] = false
	primes[1] = false
	primes[2] = true

	var i int
	for {
		for i < max && !primes[i] {
			i++
		}
		for j := i << 1; j < max; j += i {
			primes[j] = false
		}
		i++
		if i > max {
			break
		}
	}

	primesList := make([]int, 0)
	for i, p := range primes {
		if p {
			primesList = append(primesList, i)
		}
	}

	return primesList
}

var primesList = Sieve(104745)

func Nth(n int) (int, bool) {
	if n < 1 {
		return 0, false
	}
	return primesList[n-1], true
}
