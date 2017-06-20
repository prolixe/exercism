package lsproduct

import (
	"errors"
	"strconv"
)

const testVersion = 5

func LargestSeriesProduct(digits string, span int) (int, error) {

	if span > len(digits) || span < 0 {
		return -1, errors.New("Invalid")
	}
	largest := 0
	for i := 0; i <= len(digits)-span; i++ {
		product, err := serieProduct(digits[i : i+span])
		if err != nil {
			return -1, err
		}
		if product > largest {
			largest = product
		}
	}
	return largest, nil
}

func serieProduct(digits string) (product int, err error) {
	product = 1
	var d int
	for _, digit := range digits {
		if d, err = strconv.Atoi(string(digit)); err == nil {
			product *= d
		} else {
			return
		}
	}
	return
}
