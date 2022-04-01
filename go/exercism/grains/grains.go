// Package grains does calculations for grains of wheat on a chess board
package grains

import (
	"errors"
	"math"
)

// Square returns the number of grains on the current square
func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, errors.New("Oh boy...")
	}
	return uint64(math.Pow(2, float64(number-1))), nil
}

// Total sums all the grains on all squares
func Total() uint64 {
	total := uint64(0)
	square := 1

	for square < 65 {
		curr, _ := Square(square)
		total += curr
		square++
	}

	return uint64(total)
}
