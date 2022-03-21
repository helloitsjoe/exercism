package hamming

import "testing"

// type testCase struct {
// 	expectError bool
// 	want        int
// 	s1          string
// 	s2          string
// }

// var testCases = []testCase{
// 	{false, 1, "AAA", "AAT"},
// 	{false, 2, "ATA", "AAT"},
// 	{true, 1, "AAA", "AATA"},
// }

func TestHamming(t *testing.T) {
	for _, tc := range testCases {
		got, err := Distance(tc.s1, tc.s2)
		if tc.expectError {
			// check if err is of error type
			var _ error = err
			// we expect error
			if err == nil {
				t.Fatalf("Distance(%q, %q); expected error, got nil.",
					tc.s1, tc.s2)
			}
		} else {
			// we do not expect error
			if err != nil {
				t.Fatalf("Distance(%q, %q) returned unexpected error: %v",
					tc.s1, tc.s2, err)
			}
			if got != tc.want {
				t.Fatalf("Distance(%q, %q) = %d, want %d.",
					tc.s1, tc.s2, got, tc.want)
			}
		}
	}
}
func BenchmarkHamming(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	// bench combined time to run through all test cases
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			// ignoring errors and results because we're just timing function execution
			_, _ = Distance(tc.s1, tc.s2)
		}
	}
}

// Source: exercism/problem-specifications
// Commit: 4119671 Hamming: Add a tests to avoid wrong recursion solution (#1450)
// Problem Specifications Version: 2.3.0

var testCases = []struct {
	s1          string
	s2          string
	want        int
	expectError bool
}{
	{ // empty strands
		"",
		"",
		0,
		false,
	},
	{ // single letter identical strands
		"A",
		"A",
		0,
		false,
	},
	{ // single letter different strands
		"G",
		"T",
		1,
		false,
	},
	{ // long identical strands
		"GGACTGAAATCTG",
		"GGACTGAAATCTG",
		0,
		false,
	},
	{ // long different strands
		"GGACGGATTCTG",
		"AGGACGGATTCT",
		9,
		false,
	},
	{ // disallow first strand longer
		"AATG",
		"AAA",
		0,
		true,
	},
	{ // disallow second strand longer
		"ATA",
		"AGTG",
		0,
		true,
	},
	{ // disallow left empty strand
		"",
		"G",
		0,
		true,
	},
	{ // disallow right empty strand
		"G",
		"",
		0,
		true,
	},
}
