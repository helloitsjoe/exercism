// Package proverb creates a proverb from a list of items
package proverb

import "fmt"

// Proverb does the converting
func Proverb(rhyme []string) []string {
	output := []string{}

	if len(rhyme) == 0 {
		return output
	}

	head := rhyme[0]
	tail := rhyme[1:]
	prev := head

	for _, curr := range tail {
		newLine := fmt.Sprintf("For want of a %s the %s was lost.", prev, curr)
		output = append(output, newLine)
		prev = curr
	}

	lastLine := fmt.Sprintf("And all for the want of a %s.", head)
	output = append(output, lastLine)

	return output
}
