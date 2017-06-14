package grains

import "errors"

const testVersion = 1

const maxUint64 = (1 << 64) - 1

func Total() uint64 {
	return uint64(maxUint64)
}

func Square(n int) (uint64, error) {
	if 1 > n || n > 64 {
		return uint64(0), errors.New("Invalid")
	}

	return 1 << uint64(n-1), nil
}
