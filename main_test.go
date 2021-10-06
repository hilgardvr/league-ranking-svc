package main

import (
	"os"
	"reflect"
	"testing"
)

func TestSortTable(t *testing.T) {
	unsortedTable := map[string]int{
		"Tarantulas": 6,
		"Snakes":     1,
		"FC Awesome": 1,
		"Lions":      5,
		"Grouches":   0,
	}
	sorted := []tableTeam{
		{
			teamName: "Tarantulas",
			points:   6,
		}, {
			teamName: "Lions",
			points:   5,
		}, {
			teamName: "FC Awesome",
			points:   1,
		}, {
			teamName: "Snakes",
			points:   1,
		}, {
			teamName: "Grouches",
			points:   0,
		},
	}
	tableUnderTest := sortTable(unsortedTable)
	if len(tableUnderTest) != len(sorted) {
		t.Fatalf("Sorted table incorrect length")
	}
	for i, v := range sorted {
		if v != tableUnderTest[i] {
			t.Fatalf("Sorted table not in correct order")
		}
	}
}

func TestGetTeamScore(t *testing.T) {
	line := "  Liverpool   10  "
	team, score := getTeamScore(line)
	if team != "Liverpool" {
		t.Fatalf("Incorrect team name: `%s`", team)
	}
	if score != 10 {
		t.Fatalf("Incorrect team score: %d", score)
	}
}

func TestReadScores(t *testing.T) {
	table := map[string]int{
		"Tarantulas": 6,
		"Snakes":     1,
		"FC Awesome": 1,
		"Lions":      5,
		"Grouches":   0,
	}
	file, _ := os.Open("test_scores.txt")
	tableUnderTest := readScores(file)
	if !reflect.DeepEqual(table, tableUnderTest) {
		t.Fatalf("Table not created correctly")
	}

}
