package wordsearch

const testVersion = 3

func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {

	positions := make(map[string][2][2]int)
	for _, word := range words {
		if pos, found := solve(word, puzzle); found {
			positions[word] = pos
		}

	}
	return positions, nil
}

type Direction struct {
	x, y int
}

var directions = [8]Direction{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, 1},
	{0, -1},
	{1, -1},
	{1, 0},
	{1, 1},
}

func solve(word string, puzzle []string) (pos [2][2]int, found bool) {

	for y, row := range puzzle {
		for x, c := range row {
			if c == rune(word[0]) {
				for _, dir := range directions {
					if posLast, ok := match(word, x, y, dir, puzzle); ok {
						pos = [2][2]int{{x, y}, posLast}
						found = ok
						return
					}
				}
			}

		}
	}

	return
}

func match(word string, x, y int, dir Direction, puzzle []string) (posLast [2]int, found bool) {
	for i, c := range word {
		r, ok := puzzleAt(puzzle, x+i*dir.x, y+i*dir.y)
		if !ok {
			return
		}
		if r != byte(c) {
			return
		}
	}
	posLast = [2]int{x + (len(word)-1)*dir.x, y + (len(word)-1)*dir.y}
	found = true
	return

}

func puzzleAt(puzzle []string, x, y int) (byte, bool) {
	if y < 0 || len(puzzle) <= y {
		return ' ', false
	}
	row := puzzle[y]
	if x < 0 || len(row) <= x {
		return ' ', false
	}
	return puzzle[y][x], true
}
