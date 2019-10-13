package spiralmatrix

// SpiralMatrix produce a 'input' size matrix filled with digits
func SpiralMatrix(size int) [][]int {

	spiral := make([][]int, size)
	// Make empty spiral
	for i := 0; i < size; i++ {
		spiral[i] = make([]int, size)
	}

	//Define the limits
	top, left, bottom, right := 0, 0, size-1, size-1
	// Start at the top left

	for i := 1; i <= size*size; {
		// Go left
		for x := left; x <= right; x++ {
			spiral[top][x] = i
			i++
		}
		top++ // Top row completed

		// Go down
		for y := top; y <= bottom; y++ {
			spiral[y][right] = i
			i++
		}
		right-- // Rightmost col filled
		// Go right
		for x := right; x >= left; x-- {
			spiral[bottom][x] = i
			i++
		}
		bottom-- // Bottom  row filled
		// Go down
		for y := bottom; y >= top; y-- {
			spiral[y][left] = i
			i++
		}
		left++ // Left most col filled

	}

	return spiral
}
