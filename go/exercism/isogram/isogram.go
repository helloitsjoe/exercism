// Package isogram checks if a string contians repeat letters
package isogram

import "strings"

var replacer = strings.NewReplacer("-", "", " ", "")

// IsIsogram does the checking
func IsIsogram(word string) bool {
	word = replacer.Replace(word)

	seen := map[rune]bool{}

	for _, letter := range strings.ToLower(word) {
		if seen[letter] {
			return false
		}
		seen[letter] = true
	}

	return true
}
