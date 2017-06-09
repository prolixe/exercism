package accumulate

const testVersion = 1

func Accumulate(input []string, op func(string) string) []string {
	strings := make([]string, len(input))
	for i, s := range input {
		strings[i] = op(s)
	}
	return strings
}
