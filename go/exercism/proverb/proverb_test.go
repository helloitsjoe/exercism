package proverb

import (
	"fmt"
	"testing"
)

func TestProverb(t *testing.T) {
	for _, test := range stringTestCases {
		actual := Proverb(test.input)
		if fmt.Sprintf("%q", actual) != fmt.Sprintf("%q", test.expected) {
			t.Fatalf("FAIL %s - Proverb test [%s]\n\texpected: [%s],\n\tactual:   [%s]",
				test.description, test.input, test.expected, actual)
		}
		t.Logf("PASS %s", test.description)
	}
}
func BenchmarkProverb(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range stringTestCases {
			Proverb(test.input)
		}
	}
}

// Source: exercism/problem-specifications
// Commit: e86e97a proverb: apply "input" policy
// Problem Specifications Version: 1.1.0

type proverbTest struct {
	description string
	input       []string
	expected    []string
}

var stringTestCases = []proverbTest{
	{
		description: "zero pieces",
		input:       []string{},
		expected:    []string{},
	},
	{
		description: "one piece",
		input:       []string{"nail"},
		expected:    []string{"And all for the want of a nail."},
	},
	{
		description: "two pieces",
		input:       []string{"nail", "shoe"},
		expected:    []string{"For want of a nail the shoe was lost.", "And all for the want of a nail."},
	},
	{
		description: "three pieces",
		input:       []string{"nail", "shoe", "horse"},
		expected:    []string{"For want of a nail the shoe was lost.", "For want of a shoe the horse was lost.", "And all for the want of a nail."},
	},
	{
		description: "full proverb",
		input:       []string{"nail", "shoe", "horse", "rider", "message", "battle", "kingdom"},
		expected:    []string{"For want of a nail the shoe was lost.", "For want of a shoe the horse was lost.", "For want of a horse the rider was lost.", "For want of a rider the message was lost.", "For want of a message the battle was lost.", "For want of a battle the kingdom was lost.", "And all for the want of a nail."},
	},
	{
		description: "four pieces modernized",
		input:       []string{"pin", "gun", "soldier", "battle"},
		expected:    []string{"For want of a pin the gun was lost.", "For want of a gun the soldier was lost.", "For want of a soldier the battle was lost.", "And all for the want of a pin."},
	},
}
