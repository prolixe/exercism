package queenattack

import "errors"

const testVersion = 2

func isNotationValid(a string) bool {
	if len(a) != 2 {
		return false
	}
	if (a[0]-'a') < 0 || 7 < (a[0]-'a') {
		return false
	}
	if (a[1]-'0') < 1 || 8 < (a[1]-'0') {
		return false
	}
	return true
}

func isCommonDiagonal(w, b string) bool {

	fileDiff := (w[0] - 'a') - (b[0] - 'a')
	rankDiff := (w[1] - '0') - (b[1] - '0')

	return fileDiff == rankDiff || fileDiff == -rankDiff
}

func CanQueenAttack(w, b string) (bool, error) {
	if w == b {
		return false, errors.New("Same square")
	}
	if !isNotationValid(w) || !isNotationValid(b) {
		return false, errors.New("Invalid notation")
	}

	if w[0] == b[0] || w[1] == b[1] || isCommonDiagonal(w, b) {

		return true, nil
	}
	return false, nil
}
