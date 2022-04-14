// Package etl does the Transform step of Extract-Transform-Load.
package etl

import "strings"

// Transform takes Scrabble scores from a legacy system and transforms them
// into a more efficient system.
func Transform(in map[int][]string) map[string]int {
	out := make(map[string]int)

	for value, stringSlice := range in {
		for _, letter := range stringSlice {
			out[strings.ToLower(letter)] = value
		}
	}

	return out
}
