// Package diffsquares does math with squares
package diffsquares

import "math"

func intSquare(n int) int {
	return int(math.Pow(float64(n), 2))
}

// SquareOfSum returns the square of the sum of natural numbers from 1 to n
func SquareOfSum(n int) int {
	sum := 0
	for n > 0 {
		sum += n
		n--
	}
	return intSquare(sum)
}

// SquareOfSum returns the sum of the squares of natural numbers from 1 to n
func SumOfSquares(n int) int {
	sum := 0
	for n > 0 {
		sum += intSquare(n)
		n--
	}
	return sum
}

// Difference returns the difference between the sum of squares and the square of the sum of natural numbers from 1 to n
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
