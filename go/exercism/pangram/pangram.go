// Package pangram determines if a sentence uses every letter of the alphabet.
package pangram

import (
	"strings"
	"unicode"
)

// IsPangram returns a bool.
func IsPangram(input string) bool {
	counter := 0
	letters := make(map[rune]bool)

	for _, letter := range strings.ToLower(input) {
		_, ok := letters[letter]
		if ok || !unicode.IsLetter(letter) {
			continue
		}
		counter++
		if counter == 26 {
			return true
		}
		letters[letter] = true
	}
	return false
}
