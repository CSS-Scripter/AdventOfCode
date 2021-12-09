package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var (
	horizontalPositions []int
)

// type FuelRequirement struct {
// 	pos          int
// 	fuelRequired int
// }

func main() {
	prepareInput("input.txt")
	two()
}

func prepareInput(file string) {
	input, _ := ioutil.ReadFile(file)
	inputNums := strings.Split(strings.TrimSpace(string(input)), ",")
	for _, inputNum := range inputNums {
		pos, _ := strconv.Atoi(inputNum)
		horizontalPositions = append(horizontalPositions, pos)
	}
}

func one() {
	NO_POS_SET := -1
	leastAmountOfFuel := 0
	leastAmountPosition := NO_POS_SET
	for i := 0; i < 1200; i++ {
		totalFuel := 0
		for _, crabPos := range horizontalPositions {
			totalFuel += abs(crabPos - i)
		}
		if totalFuel < leastAmountOfFuel || leastAmountPosition == NO_POS_SET {
			leastAmountOfFuel = totalFuel
			leastAmountPosition = i
		}
	}
	fmt.Printf("Closest position: %d, fuel required: %d", leastAmountPosition, leastAmountOfFuel)
}

func abs(in int) int {
	if in < 0 {
		return -in
	}
	return in
}

func two() {
	NO_POS_SET := -1
	leastAmountOfFuel := 0
	leastAmountPosition := NO_POS_SET
	for i := 0; i < 1200; i++ {
		totalFuel := 0
		for _, crabPos := range horizontalPositions {
			steps := abs(crabPos - i)
			if steps%2 == 0 {
				totalFuel += (steps / 2) * (steps + 1)
			} else {
				totalFuel += (((steps - 1) / 2) + 1) * (steps)
			}
		}
		if totalFuel < leastAmountOfFuel || leastAmountPosition == NO_POS_SET {
			leastAmountOfFuel = totalFuel
			leastAmountPosition = i
		}
	}
	fmt.Printf("Closest position: %d, fuel required: %d", leastAmountPosition, leastAmountOfFuel)
}
