package pascal

const testVersion = 1

func triangleRow(n int) []int {
	if n == 1 {
		return []int{1}
	}
	row := make([]int, n)
	row[0] = 1
	topRow := triangleRow(n - 1)
	for i := range topRow[:len(topRow)-1] {
		row[i+1] = topRow[i] + topRow[i+1]
	}
	row[n-1] = 1
	return row
}

func Triangle(n int) [][]int {
	triangle := make([][]int, 0)
	for i := 1; i <= n; i++ {
		triangle = append(triangle, triangleRow(i))
	}
	return triangle
}
