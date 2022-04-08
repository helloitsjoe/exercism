// Package cryptosquare encodes secret messages
package cryptosquare

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

func filterAlphanumeric(s string) []string {
	// Remove spaces and punctuation
	letters := make([]string, 0, len(s))
	for _, char := range s {
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			letters = append(letters, string(char))
		}
	}
	fmt.Println("letters", letters)
	return letters
}

func getBlockLength(s []string) (int, int) {
	// r * c > len(pt)
	// c > r
	// c - r <= 1
	sqrt := math.Sqrt(float64(len(s)))

	r := math.Floor(sqrt)
	c := math.Floor(sqrt)

	if math.Round(sqrt) == sqrt {
		fmt.Println("equal:", math.Round(sqrt), sqrt)
		return int(sqrt), int(sqrt)
	}
	// if < 0.5, col += 1
	// if > 0.5, both += 1
	if math.Round(sqrt) == math.Floor(sqrt) {
		fmt.Println("< 0.5:", math.Round(sqrt), math.Floor(sqrt))
		r++
	}

	if math.Round(sqrt) == math.Ceil(sqrt) {
		fmt.Println("> 0.5:", math.Round(sqrt), math.Ceil(sqrt))
		c++
		r++
	}

	return int(c), int(r)
}

func getBlock(s []string) [][]string {
	blockLength, rows := getBlockLength(s)
	fmt.Println("blockLength:", blockLength)

	outer := make([][]string, 0, rows)
	for i := 0; i < blockLength; i++ {
		inner := make([]string, 0, blockLength)
		for j := 0; j < rows; j++ {
			// What index to take here?
			inner = append(inner, s[j+(blockLength*i)])
		}
		outer = append(outer, inner)
	}

	fmt.Println("outer", outer)
	return outer
}

func swapColsAndRows(m [][]string) string {
	str := ""
	counter := 0

	for i := range m[0] {
		for j := range m {
			fmt.Println("str", str)
			str += m[j][i]
			counter++
			if counter == len(m) {
				str += " "
				counter = 0
			}
		}
	}
	fmt.Println("str", str)
	return strings.TrimSpace(str)
}

// Encode does the encoding
func Encode(pt string) string {
	// return columns swapped with rows
	// single sting separated by spaces
	// pad ends with single space if necessary
	// lowercase
	compact := filterAlphanumeric(strings.ToLower(pt))
	block := getBlock(compact)
	swapped := swapColsAndRows(block)

	return swapped
}
