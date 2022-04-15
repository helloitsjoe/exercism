// Package romannumerals converts numbers into Roman numerals.
package romannumerals

import (
	"errors"
	"strings"
)

var values = []int{1000, 500, 100, 50, 10, 5, 1}
var numerals = []string{"M", "D", "C", "L", "X", "V", "I"}

// ToRomanNumeral takes in an int input and returns the Roman string.
func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input > 3000 {
		return "", errors.New("Input must be a positive integer")
	}

	output := ""
	valIdx := 0

	rounds := 0

	for input > 0 {
		denominator := values[valIdx]

		if input > 89 && input < 100 {
			output += "XC"
			input -= 90
		} else if input > 899 && input < 1000 {
			output += "CM"
			input -= 900
		} else if input == 9 {
			output += "IX"
			input = 0
		}

		repeat := input / denominator
		remainder := input % denominator

		if repeat == 4 {
			output += numerals[valIdx] + numerals[valIdx-1]
		} else {
			output += strings.Repeat(numerals[valIdx], repeat)
		}

		input = remainder

		valIdx++
		rounds++

		if rounds == 100 {
			return "", errors.New("Too many rounds")
		}
	}

	return output, nil
}
