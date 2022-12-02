package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	Rock     = 1
	Paper    = 2
	Scissors = 3

	Lose = 0
	Draw = 3
	Win  = 6

	WantWin  = 1
	WantDraw = 0
	WantLose = -1
)

var loop = []int{3, 1, 2, 3, 1}

var inputToAction = map[string]int{
	"A":  Rock,
	"B":  Paper,
	"C":  Scissors,
	"X":  Rock,
	"Y":  Paper,
	"Z":  Scissors,
	"wX": WantLose,
	"wY": WantDraw,
	"wZ": WantWin,
}

type Match struct {
	Opponent int
	You      int
}

var oneMatches []Match
var twoMatches []Match

func (m *Match) CalculateScore() int {
	if m.You == m.Opponent {
		return Draw + m.You
	}
	if loop[m.You+1] == m.Opponent {
		return Lose + m.You
	}
	return Win + m.You
}

func main() {
	prepareInput("input.txt")
	printScoreOfMatches(oneMatches)
	printScoreOfMatches(twoMatches)
}

func printScoreOfMatches(matches []Match) {
	var totalScore int
	for _, match := range matches {
		totalScore += match.CalculateScore()
	}
	fmt.Println(totalScore)
}

func prepareInput(file string) {
	input, _ := ioutil.ReadFile(file)
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	for _, line := range lines {
		actions := strings.Fields(line)

		// Prep input for one
		match := Match{
			You:      decodeAction(actions[1]),
			Opponent: decodeAction(actions[0]),
		}
		oneMatches = append(oneMatches, match)

		// Prep input for two
		opponent := decodeAction(actions[0])
		wantedOutcome := decodeAction("w" + actions[1])
		match = Match{
			You:      reverseAction(opponent, wantedOutcome),
			Opponent: opponent,
		}
		twoMatches = append(twoMatches, match)
	}
}

func decodeAction(action string) int {
	return inputToAction[action]
}

func reverseAction(action int, want int) int {
	yourAction := loop[action+want]
	return yourAction
}
