// Given a phrase, will count the occurrences of each word in that phrase A
// word will either be a collection of digits, a collection of ascii letters,
// Or a contraction of two simple words with a single apostrophe. The count is
// case-insensitive and unordered, and any punctuation besides apostrophe in a
// contraction is ignored. Words can be separated by any form of whitespace.
package wordcount

import (
	"strings"
	"unicode"
)

type Frequency map[string]int

func removeNonAlpha(word string) string {
	var alpha string

	for i, char := range word {
		isApostrophe := (char == '\'' && (i != 0 && i != len(word)-1))
		if unicode.IsLetter(char) || unicode.IsNumber(char) || isApostrophe {
			alpha += strings.ToLower(string(char))
		}
	}

	return alpha
}

// WordCount counts the words
func WordCount(phrase string) Frequency {
	f := Frequency{}
	p := strings.Replace(phrase, ",", " ", -1)

	for _, w := range strings.Fields(p) {
		word := removeNonAlpha(w)
		_, ok := f[word]
		if !ok {
			f[word] = 0
		}
		f[word] += 1
	}

	return f
}
