package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const WIN = 3
const DRAW = 1
const LOOSE = 0

type tableTeam struct {
	teamName	string
	points		int 
}

func sortTable(unsortedTable map[string]int) []tableTeam {
	table := make([]tableTeam, 0, len(unsortedTable))
	for team := range unsortedTable {
		table = append(table, tableTeam{teamName: team, points: unsortedTable[team]})
	}
	sort.Slice(table, func(i, j int) bool {
		if table[i].points == table[j].points {
			return table[i].teamName < table[j].teamName
		} else {
			return table[i].points > table[j].points
		}
	})
	return table
}

func getTeamScore(result string) (string, int) {
	trimResult := strings.TrimSpace(result)
	lastSpace := strings.LastIndex(trimResult, " ")
	teamName := trimResult[0:lastSpace]
	teamGoals := trimResult[lastSpace + 1:]
	goals, err := strconv.Atoi(string(teamGoals))
	if err != nil {
		panic(fmt.Sprintf("Could not parse teams score as int: %s", err.Error()))
	}
	return teamName, goals
}

func readScores(file *os.File) map[string]int {
	var totals = map[string]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() && scanner.Text() != "quit" {
		scores := strings.Split(scanner.Text(), ",")
		if len(scores) != 2 {
			panic(fmt.Sprintf("Line doesn't contain 2 scores: %s", scanner.Text()))
		}
		team1, score1 := getTeamScore(scores[0])
		team2, score2 := getTeamScore(scores[1])
		if score1 > score2 {
			totals[team1] += WIN
			totals[team2] += LOOSE
		} else if score1 < score2 {
			totals[team2] += WIN
			totals[team1] += LOOSE
		} else {
			totals[team1] += DRAW
			totals[team2] += DRAW
		}
	}
	return totals
}


func main() {
	args := os.Args
	var totals map[string]int
	if len(args) <= 1 {
		fmt.Println("No file specified, reading from commandline. Type `quit` to exit")
		totals = readScores(os.Stdin)
	} else {
		file, err := os.Open(args[1])
		if err != nil {
			fmt.Println("Could not open file", err)
			return
		}
		defer file.Close()
		totals = readScores(file)
	}
	sortedTotals := sortTable(totals)
	for i, v := range sortedTotals {
		fmt.Printf("%d. %s\t%d\n", i + 1, v.teamName, v.points)
	}
}