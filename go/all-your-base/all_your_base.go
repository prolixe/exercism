package allyourbase

import "errors"

const testVersion = 1

var ErrInvalidDigit = errors.New("invalid digit")
var ErrInvalidBase = errors.New("invalid base")

func ConvertToBase(base uint64, digits []uint64, outputBase uint64) ([]uint64, error) {

	if base < 2 || outputBase < 2 {
		return nil, ErrInvalidBase
	}
	for _, d := range digits {
		if d >= base {
			return nil, ErrInvalidDigit
		}
	}
	if len(digits) == 0 {
		return []uint64{0}, nil
	}
	return toDigits(fromDigits(digits, base), outputBase), nil
}

func toDigits(n, base uint64) []uint64 {

	if n == 0 {
		return []uint64{0}
	}
	digits := make([]uint64, 0)
	for n > 0 {
		d := make([]uint64, 1)
		d[0] = n % base
		// push to front
		digits = append(d, digits...)
		n /= base
	}
	return digits
}

func fromDigits(digits []uint64, base uint64) uint64 {
	var n uint64
	for _, d := range digits {
		n = n*base + d
	}
	return n
}
