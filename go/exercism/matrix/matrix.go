// Package matrix manipulates rows an columns of a matrix
package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix struct {
	rows [][]int
	cols [][]int
}

// New creates a new matrix from a string
func New(s string) (*Matrix, error) {
	var rows = [][]int{}
	var cols = [][]int{}

	for _, strrow := range strings.Split(s, "\n") {

		row := []int{}
		for _, num := range strings.Split(strings.TrimSpace(strrow), " ") {
			integer, err := strconv.ParseInt(num, 0, 64)
			if err != nil {
				return nil, errors.New("Too big for int64")
			}
			row = append(row, int(integer))
		}

		if len(rows) > 0 && len(row) != len(rows[0]) {
			return nil, errors.New("Rows are different lengths")
		}

		rows = append(rows, row)
	}

	for i := 0; i < len(rows[0]); i++ {
		col := []int{}
		for j := 0; j < len(rows); j++ {
			col = append(col, rows[j][i])
		}
		cols = append(cols, col)
	}

	return &Matrix{rows, cols}, nil
}

// Cols returns a copy of the column data
func (m *Matrix) Cols() [][]int {
	colsCopy := [][]int{}
	for _, col := range m.cols {
		colCopy := []int{}
		for _, num := range col {
			colCopy = append(colCopy, num)
		}
		colsCopy = append(colsCopy, colCopy)
	}
	return colsCopy
}

// Rows returns a copy of the row data
func (m *Matrix) Rows() [][]int {
	rowsCopy := [][]int{}
	for _, row := range m.rows {
		rowCopy := []int{}
		for _, num := range row {
			rowCopy = append(rowCopy, num)
		}
		rowsCopy = append(rowsCopy, rowCopy)
	}
	return rowsCopy
}

// Set sets a number in the matrix
func (m *Matrix) Set(row, col, val int) bool {
	if row < 0 || row >= len(m.rows) || col < 0 || col >= len(m.cols) {
		return false
	}
	m.rows[row][col] = val
	m.cols[col][row] = val
	return true
}
