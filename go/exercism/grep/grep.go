package grep

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func addNum(text string, i int) string {
	lineNum := fmt.Sprint(i + 1)
	return lineNum + ":" + text
}

func processFlags(flags []string) (bool, bool, bool, bool, bool) {
	showNums := false
	fileOnly := false
	invert := false
	insensitive := false
	entireLine := false

	for _, flag := range flags {
		switch flag {
		case "-n":
			showNums = true
		case "-l":
			fileOnly = true
		case "-i":
			insensitive = true
		case "-v":
			invert = true
		case "-x":
			entireLine = true
		}
	}

	return showNums, fileOnly, invert, insensitive, entireLine
}

func Search(pattern string, flags, files []string) []string {
	found := []string{}

	showNums, fileOnly, invert, insensitive, entireLine := processFlags(flags)

	if entireLine {
		pattern = "^" + pattern + "$"
	}
	if insensitive {
		pattern = "(?i)" + pattern
	}

	for _, file := range files {
		data, _ := os.ReadFile(file)
		contents := string(data)

		for i, line := range strings.Split(contents, "\n") {
			if len(line) == 0 {
				continue
			}

			match, _ := regexp.MatchString(pattern, line)

			if showNums {
				line = addNum(line, i)
			}

			if len(files) > 1 && !fileOnly {
				line = file + ":" + line
			}

			if fileOnly {
				if len(found) > 0 && found[len(found)-1] == file {
					continue
				}
				line = file
			}

			if invert && !match {
				found = append(found, line)
			} else if !invert && match {
				found = append(found, line)
			}
		}
	}

	return found
}
