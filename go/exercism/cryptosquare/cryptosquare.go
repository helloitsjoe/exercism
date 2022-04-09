// Package cryptosquare encodes secret messages
package cryptosquare

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

func filterAlphanumeric(s string) []string {
	letters := make([]string, 0, len(s))

	for _, char := range s {
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			letters = append(letters, string(char))
		}
	}

	return letters
}

func getColsRows(s []string) (int, int) {
	// Rules about rows:
	// r * c > len(s)
	// c > r
	// c - r <= 1

	sqrt := math.Sqrt(float64(len(s)))
	ceil := math.Ceil(sqrt)
	floor := math.Floor(sqrt)
	round := math.Round(sqrt)

	r := floor
	c := floor

	if round == sqrt {
		fmt.Println("equal:", round, sqrt)
		return int(sqrt), int(sqrt)
	}

	if round == floor {
		fmt.Println("< 0.5:", round, floor)
		r++
	}

	if round == ceil {
		fmt.Println("> 0.5:", round, ceil)
		c++
		r++
	}

	return int(c), int(r)
}

func getBlock(s []string) [][]string {
	cols, rows := getColsRows(s)

	// 12345678 causes slice bounds out of range error, add capacity
	// I'm not sure exactly why this happens...
	if cap(s) < cols*rows {
		s = append(s, "")
	}

	outer := make([][]string, 0, rows)
	for i := 0; i < cols; i++ {
		inner := s[i*rows : i*rows+rows]

		for i := range inner {
			if inner[i] == "" {
				fmt.Println("padding extra space")
				inner[i] = " "
			}
		}
		outer = append(outer, inner)
	}

	return outer
}

func swapColsAndRows(m [][]string) string {
	str := ""
	counter := 0

	for i := range m[0] {
		for j := range m {
			str += m[j][i]
			counter++
			if counter == len(m) {
				str += " "
				counter = 0
			}
		}
	}
	return str[:len(str)-1]
}

// Encode does the encoding
func Encode(pt string) string {
	compact := filterAlphanumeric(strings.ToLower(pt))

	if len(compact) == 0 {
		return ""
	}

	block := getBlock(compact)
	swapped := swapColsAndRows(block)

	return swapped
}
