package tournament

import (
	"encoding/csv"
	"fmt"
	"io"
	"sort"
)

const (
	outputLineLength = 56
	header           = "Team                           | MP |  W |  D |  L |  P\n"
)

type teamStatsStruct struct {
	name    string
	matches int
	wins    int
	losses  int
	draws   int
	points  int
}

var allowedMatchResult = map[string]bool{
	"win":  true,
	"loss": true,
	"draw": true,
}

// Tally tallies
func Tally(r io.Reader, w io.Writer) error {
	csvReader := csv.NewReader(r)
	csvReader.Comma = ';'
	csvReader.FieldsPerRecord = 3
	csvReader.Comment = '#'
	// Read all records from CSV
	records, err := csvReader.ReadAll()
	if err != nil {
		return fmt.Errorf("formatting error")
	}

	// Stores team-name: team-stats
	var teamStatMap map[string]*teamStatsStruct = make(map[string]*teamStatsStruct)

	//
	for _, record := range records {
		if len(record) != 3 {
			return fmt.Errorf("record must be length 3")
		} else if exists := allowedMatchResult[record[2]]; !exists {
			return fmt.Errorf("match result not allowed")
		}

		if _, exists := teamStatMap[record[0]]; exists {
			teamStatMap[record[0]].matches++
			switch record[2] {
			case "win":
				teamStatMap[record[0]].wins++
			case "loss":
				teamStatMap[record[0]].losses++
			case "draw":
				teamStatMap[record[0]].draws++
			}
		} else {
			switch record[2] {
			case "win":
				teamStatMap[record[0]] = &teamStatsStruct{name: record[0], matches: 1, wins: 1}
			case "loss":
				teamStatMap[record[0]] = &teamStatsStruct{name: record[0], matches: 1, losses: 1}
			case "draw":
				teamStatMap[record[0]] = &teamStatsStruct{name: record[0], matches: 1, draws: 1}
			}
		}

		if _, exists := teamStatMap[record[1]]; exists {
			teamStatMap[record[1]].matches++
			switch record[2] {
			case "win":
				teamStatMap[record[1]].losses++
			case "loss":
				teamStatMap[record[1]].wins++
			case "draw":
				teamStatMap[record[1]].draws++
			}
		} else {
			switch record[2] {
			case "win":
				teamStatMap[record[1]] = &teamStatsStruct{name: record[1], matches: 1, losses: 1}
			case "loss":
				teamStatMap[record[1]] = &teamStatsStruct{name: record[1], matches: 1, wins: 1}
			case "draw":
				teamStatMap[record[1]] = &teamStatsStruct{name: record[1], matches: 1, draws: 1}
			}
		}

	}

	teamStatsStructList := make([]teamStatsStruct, len(teamStatMap))
	// iterate through teams
	idx := 0
	for teamName, teamStats := range teamStatMap {
		teamStatMap[teamName].points = calculatePoints(teamStats)
		// read teamStatMap into a struct
		teamStatsStructList[idx] = *teamStats
		idx++
	}

	// Soft list of teamStatsStruct by points/name
	sort.Slice(teamStatsStructList, func(i, j int) bool {
		if teamStatsStructList[i].points > teamStatsStructList[j].points {
			return true
		}
		if teamStatsStructList[i].points < teamStatsStructList[j].points {
			return false
		}
		return teamStatsStructList[i].name < teamStatsStructList[j].name
	})

	reservedByteString := make([]byte, outputLineLength*(len(teamStatMap)+1))
	idx = 0
	// place the header
	for _, letter := range []byte(header) {
		reservedByteString[idx] = byte(letter)
		idx++
	}
	// write each line of team output
	for _, teamStats := range teamStatsStructList {
		for _, letter := range []byte(teamStats.String()) {
			reservedByteString[idx] = byte(letter)
			idx++
		}
		reservedByteString[idx] = byte('\n')
		idx++
	}

	w.Write(reservedByteString)

	return nil
}

func calculatePoints(stats *teamStatsStruct) int {
	return stats.wins*3 + stats.draws
}

func (statStruct teamStatsStruct) String() string {
	return fmt.Sprintf("%-30v | %2v | %2v | %2v | %2v | %2v", statStruct.name, statStruct.matches, statStruct.wins, statStruct.draws, statStruct.losses, statStruct.points)
}
