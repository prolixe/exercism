package minesweeper

import "regexp"

const testVersion = 1

type BadBoardError struct {
	err string
}

var validRow = regexp.MustCompile(`^\|[1-8\* ]*\|$`)

func (b *BadBoardError) Error() string {
	return b.err
}

func (b *Board) Count() error {
	if err := b.validate(); err != nil {
		return err
	}

	for i, row := range *b {
		for j := range row {
			if (*b)[i][j] == '*' {
				b.increment(i, j)
			}
		}
	}

	return nil
}

func (b *Board) increment(i, j int) {

	for x := i - 1; x <= i+1; x++ {
		for y := j - 1; y <= j+1; y++ {
			if y == j && x == i {
				continue
			}
			if (*b)[x][y] == ' ' {
				(*b)[x][y] = '1'
			} else if (byte(0) < (*b)[x][y]-'0') && ((*b)[x][y]-'0' < byte(9)) {
				(*b)[x][y] += byte(1)
			}
		}
	}

}

func (b Board) validate() error {

	if len(b) < 2 {
		return &BadBoardError{"bad dimensions"}
	}

	// Check first and last row, their content must start and end with a +, and have '-' inbetween
	firstRow := b[0]
	rowLen := len(firstRow)
	lastRow := b[len(b)-1]
	if firstRow[0] != firstRow[len(firstRow)-1] {
		return &BadBoardError{"invalid board"}
	}
	if string(lastRow) != string(firstRow) {
		return &BadBoardError{"invalid board"}
	}
	// Check the sides, must start and end with '|'
	// and all be the same length
	if len(b) > 2 {
		for _, row := range b[1 : len(b)-1] {
			// A row must only contain digits, whitespace or '*'
			if !validRow.Match(row) || len(row) != rowLen {
				return &BadBoardError{"invalid board"}
			}
		}
	}

	return nil
}
