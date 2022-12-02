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
)

type Match struct {
	Opponent int
	You      int
}

var matches []Match

func (m *Match) CalculateScore() int {
	switch true {
	case
		(m.You == Rock && m.Opponent == Rock),
		(m.You == Paper && m.Opponent == Paper),
		(m.You == Scissors && m.Opponent == Scissors):
		return Draw + m.You
	case
		(m.You == Rock && m.Opponent == Paper),
		(m.You == Paper && m.Opponent == Scissors),
		(m.You == Scissors && m.Opponent == Rock):
		return Lose + m.You
	case
		(m.You == Rock && m.Opponent == Scissors),
		(m.You == Paper && m.Opponent == Rock),
		(m.You == Scissors && m.Opponent == Paper):
		return Win + m.You
	}
	panic("uncaught combination")
}

func main() {
	prepareInput("input.txt")
	one()
	matches = []Match{}
	prepareInputTwo("input.txt")
	one()
}

func one() {
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
		actions := strings.Split(line, " ")
		opponent := actions[0]
		you := actions[1]
		match := Match{
			You:      decodeAction(you),
			Opponent: decodeAction(opponent),
		}
		matches = append(matches, match)
	}
}

func decodeAction(action string) int {
	switch action {
	case "A", "X":
		return Rock
	case "B", "Y":
		return Paper
	case "C", "Z":
		return Scissors
	default:
		panic("Unknown action: " + action)
	}
}

func prepareInputTwo(file string) {
	input, _ := ioutil.ReadFile(file)
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	for _, line := range lines {
		actions := strings.Split(line, " ")
		opponent := actions[0]
		action := actions[1]
		youAction, opponentAction := decodeActionTwo(action, opponent)
		match := Match{
			You:      youAction,
			Opponent: opponentAction,
		}
		matches = append(matches, match)
	}
}

func decodeActionTwo(action string, opponent string) (int, int) {
	switch opponent + action {
	case "AX":
		return Scissors, Rock
	case "BX":
		return Rock, Paper
	case "CX":
		return Paper, Scissors
	case "AY":
		return Rock, Rock
	case "BY":
		return Paper, Paper
	case "CY":
		return Scissors, Scissors
	case "AZ":
		return Paper, Rock
	case "BZ":
		return Scissors, Paper
	case "CZ":
		return Rock, Scissors
	default:
		panic("unknown combination")
	}
}
