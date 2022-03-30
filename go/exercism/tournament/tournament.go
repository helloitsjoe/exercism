// Package tournament tallies tournament scores
package tournament

import (
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type formattedRecord struct {
	name    string
	matches int
	win     int
	draw    int
	loss    int
	points  int
}

type record = map[string]int

var validOutcomes = map[string]bool{
	"win":  true,
	"loss": true,
	"draw": true,
}

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

func convertToWrite(s string) ([]byte, error) {
	b := make([]byte, 0)
	for _, char := range s {
		b = append(b, byte(char))
	}
	return b, nil
}

func getRecords(s string) (map[string]record, error) {
	games := strings.Split(s, "\n")
	records := make(map[string]record)

	for _, game := range games {
		if game == "" || game == " " || game[0] == '#' {
			continue
		}

		r := strings.Split(game, ";")

		if len(r) != 3 {
			return nil, errors.New("Game should have two teams and an outcome")
		}

		firstTeam := r[0]
		secondTeam := r[1]
		firstOutcome := r[2]

		if _, ok := validOutcomes[firstOutcome]; !ok {
			return nil, errors.New("Unrecognized outcome")
		}

		var secondOutcome string
		var firstPoints int
		var secondPoints int

		if firstOutcome == "win" {
			firstPoints = 3
			secondPoints = 0
			secondOutcome = "loss"
		} else if firstOutcome == "loss" {
			firstPoints = 0
			secondPoints = 3
			secondOutcome = "win"
		} else {
			firstPoints = 1
			secondPoints = 1
			secondOutcome = "draw"
		}

		var firstRecord record
		if val, ok := records[firstTeam]; ok {
			firstRecord = val
			firstRecord["matches"] += 1
			firstRecord[firstOutcome] += 1
			firstRecord["points"] += firstPoints
		} else {
			firstRecord = make(record)
			firstRecord["matches"] = 1
			firstRecord[firstOutcome] = 1
			firstRecord["points"] = firstPoints
		}

		var secondRecord record
		if val, ok := records[secondTeam]; ok {
			secondRecord = val
			secondRecord["matches"] += 1
			secondRecord[secondOutcome] += 1
			secondRecord["points"] += secondPoints
		} else {
			secondRecord = make(record)
			secondRecord["matches"] = 1
			secondRecord[secondOutcome] = 1
			secondRecord["points"] = secondPoints
		}

		records[firstTeam] = firstRecord
		records[secondTeam] = secondRecord
	}

	return records, nil
}

func formatTable(r map[string]record) string {
	toSort := []formattedRecord{}
	for k, v := range r {
		team := formattedRecord{k, v["matches"], v["win"], v["draw"], v["loss"], v["points"]}
		toSort = append(toSort, team)
	}

	sort.SliceStable(toSort, func(i, j int) bool {
		firstPoints := int(toSort[i].points)
		secondPoints := int(toSort[j].points)

		if firstPoints == secondPoints {
			return toSort[i].name < toSort[j].name
		}

		return firstPoints > secondPoints
	})

	teams := []string{fmt.Sprintf("%-30s | %2s | %2s | %2s | %2s | %2s\n", "Team", "MP", "W", "D", "L", "P")}
	for _, v := range toSort {
		team := fmt.Sprintf("%-30s | %2d | %2d | %2d | %2d | %2d\n", v.name, v.matches, v.win, v.draw, v.loss, v.points)
		teams = append(teams, team)
	}

	return strings.Join(teams, "")
}

// Tally does the tallying
func Tally(reader io.Reader, writer io.Writer) error {
	s, err := parseReader(reader)
	if err != nil {
		return err
	}

	records, err := getRecords(s)

	if err != nil {
		return err
	}

	output := formatTable(records)
	fmt.Println("output:", output)
	bytes, err := convertToWrite(output)

	if err != nil {
		return err
	}

	writer.Write(bytes)
	return nil
}
