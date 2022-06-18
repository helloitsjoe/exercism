package ocr

import "strings"

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

func combineLines(input string) []string {
	// Slice off first newline
	split := strings.Split(input, "\n")[1:]
	if len(split) == 4 {
		return []string{input}
	}

	output := []string{}
	for i := 0; i < len(split)/4; i += 1 {
		j := i * 4
		if j+4 <= len(split) {
			single := strings.Join(split[j:j+4], "\n")
			output = append(output, single)
		}
	}
	return output
}

func getDigits(input string) []string {
	split := strings.Split(input, "\n")
	digits := []string{}

	for i := 0; i < len(split[1])/3; i += 1 {
		digit := ""
		for _, line := range split {
			j := i * 3
			if j+3 <= len(line) {
				digit += "\n" + line[j:j+3]
			}
		}

		if digit != "" {
			digits = append(digits, digit)
		}
	}

	return digits
}

func Recognize(input string) []string {
	singleLines := combineLines(input)
	output := make([]string, len(singleLines))

	for i, line := range singleLines {
		for _, digit := range getDigits(line) {
			output[i] += recognizeDigit(digit)
		}
	}

	return output
}
