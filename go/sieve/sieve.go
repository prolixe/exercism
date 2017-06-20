package sieve

const testVersion = 1

// Sieve return an slice of size max containing
// a list of primes up to max -1
// This time inspired by user  xbxbxb007
func Sieve(max int) []int {
	primes := make([]int, 0)

	primes = append(primes, 2)
	for i := 3; i < max; i += 2 {
		isPrime := true
		for _, p := range primes {
			if i%p == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, i)
		}
	}

	return primes
}
