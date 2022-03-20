// Package twofer returns a very fun string.
package twofer

import "fmt"

// ShareWith does the returning.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
