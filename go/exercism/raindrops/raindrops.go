// Package raindrops is like FizzBuzz
package raindrops

import "fmt"

// Convert converts a number into raindrop sounds
func Convert(number int) string {
	// If the number has 3 as a factor, add "Pling"
	// If the number has 5 as a factor, add "Plang"
	// If the number has 7 as a factor, add "Plong"
	// Otherwise return the number

	result := ""

	if number%3 == 0 {
		result += "Pling"
	}

	if number%5 == 0 {
		result += "Plang"
	}

	if number%7 == 0 {
		result += "Plong"
	}

	if result == "" {
		result = fmt.Sprintf("%d", number)
	}

	return result
}
