package palindrome

import (
	"errors"
	"strconv"
)

const testVersion = 1

type Product struct {
	product        int
	Factorizations [][2]int
}

func Products(min, max int) (pmin, pmax Product, err error) {

	if max < min {
		err = errors.New("fmin > fmax")
		return
	}

	for i := min; i <= max; i++ {
		for j := min; j <= i; j++ {
			p := Product{i * j, nil}
			if isPalindrome(p.product) {
				if p.product < pmin.product {
					pmin = p
					pmin.Factorizations = append(pmin.Factorizations, [2]int{j, i})
				} else if p.product == pmin.product {
					pmin.Factorizations = append(pmin.Factorizations, [2]int{j, i})
				}
				if pmax.product < p.product {
					pmax = p
					pmax.Factorizations = append(pmax.Factorizations, [2]int{j, i})
				} else if p.product == pmax.product {
					pmax.Factorizations = append(pmax.Factorizations, [2]int{j, i})
				}

			}
		}
	}
	if pmin.product == 0 && pmax.product == 0 {
		err = errors.New("no palindromes")
	}

	return
}

func isPalindrome(n int) bool {
	// Trim any negative sign
	if n < 0 {
		n = -n
	}
	s := strconv.Itoa(n)
	for i := 0; i < len(s)/2+1; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
