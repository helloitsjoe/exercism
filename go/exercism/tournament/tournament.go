// Package tournament tallies tournament scores
package tournament

import (
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type recordSlice struct {
	name    string
	matches int
	win     int
	draw    int
	loss    int
	points  int
}

type record = map[string]int

var validOutcomes = map[string]bool{"win": true, "loss": true, "draw": true}

var pointsMap = map[string]int{
	"win":  3,
	"draw": 1,
	"loss": 0,
}

var outcomesMap = map[string]string{
	"win":  "loss",
	"loss": "win",
	"draw": "draw",
}

func parseReader(reader io.Reader) (string, error) {
	b := make([]byte, 8)
	s := ""
	for {
		n, err := reader.Read(b)
		if err != nil {
			if err == io.EOF {
				break
			}

			return "", err
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

func updateRecords(records map[string]record, team string, outcome string) {
	record := records[team]

	if record == nil {
		record = make(map[string]int)
		record[outcome] = 0
		record["matches"] = 0
		record["points"] = 0
	}

	record[outcome] += 1
	record["matches"] += 1
	record["points"] += pointsMap[outcome]

	records[team] = record
}

func getRecords(s string) (map[string]record, error) {
	records := make(map[string]record)

	for _, game := range strings.Split(s, "\n") {
		if strings.TrimSpace(game) == "" || game[0] == '#' {
			continue
		}

		r := strings.Split(game, ";")
		if len(r) != 3 {
			return nil, errors.New("Game should have two teams and an outcome")
		}

		firstTeam := r[0]
		secondTeam := r[1]
		firstOutcome := r[2]
		secondOutcome := outcomesMap[firstOutcome]

		if !validOutcomes[firstOutcome] || !validOutcomes[secondOutcome] {
			return nil, errors.New("Unrecognized outcome")
		}

		updateRecords(records, firstTeam, firstOutcome)
		updateRecords(records, secondTeam, secondOutcome)
	}

	return records, nil
}

func formatTable(r map[string]record) string {
	toSort := []recordSlice{}
	for k, v := range r {
		team := recordSlice{k, v["matches"], v["win"], v["draw"], v["loss"], v["points"]}
		toSort = append(toSort, team)
	}

	sort.SliceStable(toSort, func(i, j int) bool {
		firstPoints := toSort[i].points
		secondPoints := toSort[j].points

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

// Tally takes unformatted matches from a reader, formats them into a table,
// and writes them out to a writer
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

	bytes, err := convertToWrite(output)
	if err != nil {
		return err
	}

	writer.Write(bytes)
	return nil
}
