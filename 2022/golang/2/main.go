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

var loop = []int{3, 1, 2, 3, 1}
var outcomeList = []int{Win, Lose, Draw, Win, Lose}
var oneMatches []Match
var twoMatches []Match

type Match struct {
	Opponent int
	You      int
}

func (m *Match) CalculateScore() int {
	return outcomeList[m.You-m.Opponent+2] + m.You
}

func main() {
	prepareInput("input.txt")
	printScoreOfMatches(oneMatches)
	printScoreOfMatches(twoMatches)
}

func prepareInput(file string) {
	input, _ := ioutil.ReadFile(file)
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	for _, line := range lines {
		actions := strings.Fields(line)

		// Prep input for one
		match := Match{
			You:      inputToAction[actions[1]],
			Opponent: inputToAction[actions[0]],
		}
		oneMatches = append(oneMatches, match)

		// Prep input for two
		opponent := inputToAction[actions[0]]
		wantedOutcome := inputToAction["w"+actions[1]]
		match = Match{
			You:      loop[opponent+wantedOutcome],
			Opponent: opponent,
		}
		twoMatches = append(twoMatches, match)
	}
}

func printScoreOfMatches(matches []Match) {
	var totalScore int
	for _, match := range matches {
		totalScore += match.CalculateScore()
	}
	fmt.Println(totalScore)
}
