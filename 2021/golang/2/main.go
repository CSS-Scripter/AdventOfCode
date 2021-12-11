package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var (
	horizontal = 0
	depth      = 0
	aim        = 0
	movements  []Movement
)

type Movement struct {
	Dir    string
	Amount int
}

func (m Movement) executeOne() {
	if m.Dir == "forward" {
		horizontal += m.Amount
		return
	}
	if m.Dir == "down" {
		depth += m.Amount
		return
	}
	if m.Dir == "up" {
		depth -= m.Amount
		return
	}
}

func (m Movement) executeTwo() {
	if m.Dir == "forward" {
		horizontal += m.Amount
		depth += aim * m.Amount
		return
	}
	if m.Dir == "down" {
		aim += m.Amount
		return
	}
	if m.Dir == "up" {
		aim -= m.Amount
		return
	}
}

func main() {
	one()
	two()
}

func one() {
	prepareInput("input.txt")
	for _, mov := range movements {
		mov.executeOne()
	}
	fmt.Printf("One: %d\n", depth*horizontal)
}

func two() {
	prepareInput("input.txt")
	for _, mov := range movements {
		mov.executeTwo()
	}
	fmt.Printf("One: %d\n", depth*horizontal)
}

func prepareInput(file string) {
	horizontal = 0
	depth = 0
	aim = 0
	movements = []Movement{}
	input, _ := ioutil.ReadFile(file)
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	for _, line := range lines {
		dir := strings.Split(line, " ")[0]
		amount, _ := strconv.Atoi(strings.Split(line, " ")[1])
		movement := Movement{Dir: dir, Amount: amount}
		movements = append(movements, movement)
	}
}
