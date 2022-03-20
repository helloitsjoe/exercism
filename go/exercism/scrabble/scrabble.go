// Package scrabble scores a Scrabble word
package scrabble

import (
	"fmt"
	"strings"
)

// makeScoreMap creates a map of letters to their values
func makeScoreMap() map[string]int {
	scoreMap := map[int]string{
		1:  "AEIOULNRST",
		2:  "DG",
		3:  "BCMP",
		4:  "FHVWY",
		5:  "K",
		8:  "JX",
		10: "QZ",
	}

	returnMap := map[string]int{}

	for k, v := range scoreMap {
		for _, char := range v {
			returnMap[string(char)] = k
		}
	}

	return returnMap
}

// Score returns the score for a given word
func Score(word string) int {
	fmt.Println(word)
	scoreMap := makeScoreMap()

	score := 0

	for _, letter := range strings.ToUpper(word) {
		score += scoreMap[string(letter)]
	}

	return score
}
