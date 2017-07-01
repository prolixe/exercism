package perfect

const testVersion = 1

var ErrOnlyPositive error

type Classification int

const (
	ClassificationInvalid Classification = iota
	ClassificationDeficient
	ClassificationPerfect
	ClassificationAbundant
)

func Classify(n uint64) (Classification, error) {

	if n < 1 {
		return ClassificationInvalid, ErrOnlyPositive
	}
	if n == 1 {
		return ClassificationDeficient, nil
	}

	sum := uint64(1)
	// No need to go over n/2 + 1 since above
	// that we cannot divide by an integer that is not n itself
	for i := uint64(2); i < n/2+1; i++ {
		if n%i == 0 {
			sum += i
		}
	}

	if sum < n {
		return ClassificationDeficient, nil
	} else if sum > n {
		return ClassificationAbundant, nil
	}
	return ClassificationPerfect, nil
}
