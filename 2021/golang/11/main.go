package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var (
	squids Squids
)

type Squid struct {
	GlowLevel  int
	HasFlashed bool
}

type Squids [][]Squid

func (s Squids) String() string {
	returnString := ""
	for _, row := range s {
		returnString += fmt.Sprintf("%v\n", row)
	}
	return returnString
}

func (s Squid) String() string {
	return fmt.Sprint(s.GlowLevel)
}

func main() {
	one()
}

func one() {
	prepareInput("input.txt")
	simulateSteps(500)
}

func prepareInput(file string) {
	input, _ := ioutil.ReadFile(file)
	inputStrings := strings.Split(strings.TrimSpace(string(input)), "\n")
	for _, inputLine := range inputStrings {
		row := []Squid{}
		for _, squidLevel := range inputLine {
			glowLevel, _ := strconv.Atoi(string(squidLevel))
			squid := Squid{GlowLevel: glowLevel, HasFlashed: false}
			row = append(row, squid)
		}
		squids = append(squids, row)
	}
}

func simulateSteps(steps int) {
	flashes := 0
	for i := 0; i < steps; i++ {
		increaseGlowLevels()
		flashes += flash()
		if checkSynchronisation() {
			fmt.Printf("Two: %d\n", i+1)
		}
		resetFlashes()
		// fmt.Println(squids)
	}
	fmt.Printf("One: %d\n", flashes)
}

func increaseGlowLevels() {
	for i, row := range squids {
		for j := range row {
			squids[i][j].GlowLevel += 1
		}
	}
}

func flash() int {
	flashes := 0
	for y, row := range squids {
		for x := range row {
			if squids[x][y].GlowLevel >= 10 && squids[x][y].HasFlashed == false {
				squids[x][y].HasFlashed = true
				flashes += 1
				increaseGlowAround(x, y)
			}
		}
	}
	if flashes > 0 {
		flashes += flash()
	}
	return flashes
}

func increaseGlowAround(x int, y int) {
	if !isOutOfBounds(x-1, y-1) {
		squids[x-1][y-1].GlowLevel += 1
	}
	if !isOutOfBounds(x-1, y) {
		squids[x-1][y].GlowLevel += 1
	}
	if !isOutOfBounds(x-1, y+1) {
		squids[x-1][y+1].GlowLevel += 1
	}
	if !isOutOfBounds(x, y-1) {
		squids[x][y-1].GlowLevel += 1
	}
	if !isOutOfBounds(x, y+1) {
		squids[x][y+1].GlowLevel += 1
	}
	if !isOutOfBounds(x+1, y-1) {
		squids[x+1][y-1].GlowLevel += 1
	}
	if !isOutOfBounds(x+1, y) {
		squids[x+1][y].GlowLevel += 1
	}
	if !isOutOfBounds(x+1, y+1) {
		squids[x+1][y+1].GlowLevel += 1
	}
}

func checkSynchronisation() bool {
	for y, row := range squids {
		for x := range row {
			if !squids[x][y].HasFlashed {
				return false
			}
		}
	}
	return true
}

func resetFlashes() {
	for y, row := range squids {
		for x := range row {
			if squids[x][y].GlowLevel >= 10 {
				squids[x][y].GlowLevel = 0
				squids[x][y].HasFlashed = false
			}
		}
	}
}

func isOutOfBounds(x int, y int) bool {
	return x < 0 || x >= len(squids) || y < 0 || y >= len(squids[0])
}
