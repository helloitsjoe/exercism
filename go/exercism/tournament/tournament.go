// Package tournament tallies tournament scores
package tournament

import (
	"fmt"
	"io"
	"strings"
)

// type record struct {
// 	matches int
// 	wins    int
// 	draws   int
// 	losses  int
// 	poins   int
// }

type record = map[string]int

func parseReader(reader io.Reader) (string, error) {
	b := make([]byte, 8)
	s := ""
	for {
		n, err := reader.Read(b)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return "", err
			}
		}
		s += string(b[:n])
	}
	return s, nil
}

func getRecords(s string) map[string]record {
	games := strings.Split(s, "\n")
	records := make(map[string]record)
	fmt.Println(games)
	fmt.Println(len(games))

	for _, game := range games {
		// Why is this empty sometimes?
		if game == "" {
			continue
		}

		r := strings.Split(game, ";")

		firstTeam := r[0]
		secondTeam := r[1]
		outcome := r[2]

		var oppositeOutcome string
		if outcome == "win" {
			oppositeOutcome = "loss"
		} else if outcome == "loss" {
			oppositeOutcome = "win"
		} else {
			oppositeOutcome = "draw"
		}

		var firstRecord record
		if val, ok := records[firstTeam]; ok {
			firstRecord = val
			firstRecord["matches"] += 1
			firstRecord[outcome] += 1
		} else {
			firstRecord = make(record)
			firstRecord["matches"] = 1
			firstRecord[outcome] = 1
		}

		var secondRecord record
		if val, ok := records[secondTeam]; ok {
			secondRecord = val
			secondRecord["matches"] += 1
			secondRecord[oppositeOutcome] += 1
		} else {
			secondRecord = make(record)
			secondRecord["matches"] = 1
			secondRecord[oppositeOutcome] = 1
		}

		// if val, ok := firstRecord[outcome]; ok {
		// 	firstRecord[outcome] = val + 1
		// } else {
		// 	firstRecord[outcome] = 1
		// }

		records[firstTeam] = firstRecord
		records[secondTeam] = secondRecord
	}
	fmt.Println("records", records)

	return records
}

func formatTable(r map[string]record) string {
	// create map from string
	fmt.Println(r)
	return ""
}

// Tally does the tallying
func Tally(reader io.Reader, writer io.Writer) error {
	s, err := parseReader(reader)
	if err != nil {
		return err
	}
	records := getRecords(s)
	output := formatTable(records)

	fmt.Println(records)
	fmt.Println(output)
	return nil
}
