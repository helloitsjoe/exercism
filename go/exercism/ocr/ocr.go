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
	digits := make([][]string, len(split))

	for i, line := range split {
		digit := make([]string, 16)
		for j := 0; j < len(line); j += 1 {
			digit[j] = line[i : i+3]
		}
		digits[i] = digit
	}

	// foo := ""
	// for _, line := range digits {
	// 	foo += line + "\n"
	// }

	fmt.Println("digits", digits)
	// fmt.Println("foo", foo)

	output := ""

	for _, matrix := range digits {
		for _, s := range matrix {
			output += recognizeDigit(string(s))
		}
	}

	return []string{output}
}
