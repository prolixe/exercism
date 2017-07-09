package matrix

import "math"

const testVersion = 1

type Pair [2]int

func (m Matrix) Saddle() (sp []Pair) {
	// Find the lowest element of each column
	for c, col := range m.Cols() {

		lowest := int(math.MaxInt32)
		lowestRows := []int{}
		for r, num := range col {
			if num == lowest {
				lowestRows = append(lowestRows, r)
			}
			if num < lowest {
				lowest = num
				lowestRows = []int{r}
			}

		}

		// Find the highest element in the row where the lowest col element is located
		for _, lowestRow := range lowestRows {
			largest := lowest
			for _, num := range m.Rows()[lowestRow] {
				if num > largest {
					largest = num
				}
			}
			if largest == lowest {
				sp = append(sp, Pair{lowestRow, c})
			}

		}
	}

	return
}