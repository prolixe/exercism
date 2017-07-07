package matrix

import (
	"errors"
	"strconv"
	"strings"
)

const testVersion = 1

type Matrix [][]int

func New(s string) (m Matrix, err error) {

	rows := strings.Split(s, "\n")
	m = make(Matrix, len(rows))
	for i, r := range rows {
		cols := strings.Split(strings.Trim(r, " "), " ")
		// Make sure each row is the same length as the previous one
		if i > 0 && len(cols) != len(m[i-1]) {
			err = errors.New("Uneven rows")
			return
		}

		m[i] = make([]int, len(cols))
		for j, c := range cols {
			m[i][j], err = strconv.Atoi(c)
			if err != nil {
				return nil, err
			}
		}
	}
	return m, nil
}

func (m Matrix) Rows() [][]int {
	rows := make(Matrix, len(m))

	for i, r := range m {
		rows[i] = make([]int, len(r))
		copy(rows[i], r)
	}
	return [][]int(rows)
}

func (m Matrix) Cols() [][]int {
	cols := make(Matrix, len(m[0]))
	for j := range m[0] {
		row := make([]int, len(m))
		cols[j] = row
		for i := range m {
			cols[j][i] = m[i][j]
		}
	}
	return cols
}

func (m Matrix) Set(r, c, val int) bool {
	if r < 0 || len(m) <= r {
		return false
	}
	if c < 0 || len(m[0]) <= c {
		return false
	}
	m[r][c] = val
	return true
}
