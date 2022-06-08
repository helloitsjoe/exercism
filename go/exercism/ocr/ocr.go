package ocr

import "strings"

// "fmt"
// "strings"

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
	digits := ""

	for _, line := range strings.Split(input, "\n") {
		digit := recognizeDigit(line)
		digits += digit
	}

	return []string{digits}
}
