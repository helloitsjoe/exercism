package ocr

import (
	"fmt"
	"strings"
)

func recognizeDigit(input string) string {
	const zero = `
 _ 
| |
|_|
   `
	const one = `
   
  |
  |
   `

	const two = `
 _ 
 _|
|_ 
   `

	const three = `
 _ 
 _|
 _|
   `

	const four = `
   
|_|
  |
   `

	const five = `
 _ 
|_ 
 _|
   `

	const six = `
 _ 
|_ 
|_|
   `

	const seven = `
 _ 
  |
  |
   `

	const eight = `
 _ 
|_|
|_|
   `

	const nine = `
 _ 
|_|
 _|
   `

	mapper := make(map[string]string)
	mapper[zero] = "0"
	mapper[one] = "1"
	mapper[two] = "2"
	mapper[three] = "3"
	mapper[four] = "4"
	mapper[five] = "5"
	mapper[six] = "6"
	mapper[seven] = "7"
	mapper[eight] = "8"
	mapper[nine] = "9"

	val := mapper[input]

	if val == "" {
		return "?"
	}

	return val
}

func Recognize(input string) []string {
	// TODO: Split string into 3x4 digits
	split := strings.Split(input, "\n")
	digits := make([]string, 0)

	// TODO: Handle malformed

	for i := 0; i < len(split[1])/3; i += 1 {
		digit := ""
		for _, line := range split {
			foo := i * 3
			if foo+3 <= len(line) {
				digit += "\n" + line[foo:foo+3]
			}
		}

		if digit != "" {
			digits = append(digits, digit)
		}
	}

	fmt.Println("digits", digits)

	output := ""

	for _, digit := range digits {
		output += recognizeDigit(digit)
	}

	return []string{output}
}
