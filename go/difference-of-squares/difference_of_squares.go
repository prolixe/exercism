package diffsquares

const testVersion = 1

func SumOfSquares(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		total += i * i
	}
	return total
}

func SquareOfSums(n int) int {

	total := 0
	for i := 1; i <= n; i++ {
		total += i
	}
	return total * total
}

func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
