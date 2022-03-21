// Package hamming determines Hamming distance between DNA sequences.
package hamming

import "errors"

// Distance takes two strings and returns the Hamming distance between them.
func Distance(s1 string, s2 string) (int, error) {
	if len(s1) != len(s2) {
		return -1, errors.New("Inputs must be the same length")
	}

	diff := 0

	for i, char := range s1 {
		if rune(s2[i]) != char {
			diff++
		}
	}

	return diff, nil
}
