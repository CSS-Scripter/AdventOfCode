package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	// CHAIR is a constant for the chair character
	CHAIR = 'L'
	// GROUND is a constant for the ground character
	GROUND = '.'
	// OCCUPIED is a constact for the occupied character
	OCCUPIED = '#'
)

var layout = [][]rune{}

func main() {
	readFile()
	one()
	two()
}

// Slope keeps track of a movement towars a certain direction
type Slope struct {
	x int
	y int
}

func readFile() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := []rune{}
		for _, char := range scanner.Text() {
			row = append(row, char)
		}
		layout = append(layout, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func one() {
	var newLayout [][]rune = layout
	var changes int
	for {
		newLayout, changes = playRound(newLayout)
		if changes == 0 {
			break
		}
	}
	occupiedSpots := countOccupiedChairs(newLayout)
	fmt.Printf("Occupied spots: %v\n", occupiedSpots)
}

func playRound(chairLayout [][]rune) ([][]rune, int) {
	changes := 0
	newLayout := [][]rune{}
	for y, row := range chairLayout {
		newLayout = append(newLayout, []rune{})
		for x, place := range row {
			newPlaceChar := place
			occupations := countOccupationsSurroundingSpot(chairLayout, x, y)
			// fmt.Printf("Place %v, chairs surrounding: %v, occupations surrounding: %v \n", string(place), chairs, occupations)
			switch place {
			case CHAIR:
				if occupations == 0 {
					newPlaceChar = OCCUPIED
					changes++
				}
				break
			case OCCUPIED:
				if occupations >= 4 {
					newPlaceChar = CHAIR
					changes++
				}
				break
			}
			newLayout[y] = append(newLayout[y], newPlaceChar)
		}
	}
	return newLayout, changes
}

func countOccupationsSurroundingSpot(chairLayout [][]rune, x int, y int) int {
	yPlaces := []int{y - 1, y, y + 1}
	placesToCheck := map[int][]int{x - 1: yPlaces, x: {y - 1, y + 1}, x + 1: yPlaces}
	occupationCount := 0
	for xCheck, yChecks := range placesToCheck {
		for _, yCheck := range yChecks {
			if 0 <= xCheck && xCheck < len(chairLayout[0]) && 0 <= yCheck && yCheck < len(chairLayout) {
				if chairLayout[yCheck][xCheck] == OCCUPIED {
					occupationCount++
				}
			}
		}
	}
	return occupationCount
}

func countOccupiedChairs(chairLayout [][]rune) int {
	occupiedSpots := 0
	for _, row := range chairLayout {
		for _, place := range row {
			if place == OCCUPIED {
				occupiedSpots++
			}
		}
	}
	return occupiedSpots
}

func two() {
	var newLayout [][]rune = layout
	var changes int
	for {
		newLayout, changes = playRoundTwo(newLayout)
		if changes == 0 {
			break
		}
	}
	occupiedSpots := countOccupiedChairs(newLayout)
	fmt.Printf("Occupied spots: %v\n", occupiedSpots)
}

func playRoundTwo(chairLayout [][]rune) ([][]rune, int) {
	changes := 0
	newLayout := [][]rune{}

	for y, row := range chairLayout {
		newLayout = append(newLayout, []rune{})
		for x, place := range row {
			visibleOccupations := countOccupiedChairsInSight(chairLayout, x, y)
			newPlaceChar := place
			switch place {
			case CHAIR:
				if visibleOccupations == 0 {
					newPlaceChar = OCCUPIED
					changes++
				}
				break
			case OCCUPIED:
				if visibleOccupations >= 5 {
					newPlaceChar = CHAIR
					changes++
				}
				break
			}
			newLayout[y] = append(newLayout[y], newPlaceChar)
		}
	}
	return newLayout, changes
}

func countOccupiedChairsInSight(chairLayout [][]rune, x int, y int) int {
	slopes := []Slope{
		Slope{x: -1, y: -1},
		Slope{x: -1, y: 0},
		Slope{x: -1, y: 1},
		Slope{x: 0, y: -1},
		Slope{x: 0, y: 1},
		Slope{x: 1, y: -1},
		Slope{x: 1, y: 0},
		Slope{x: 1, y: 1},
	}
	occupiedPlacesInSight := 0
	for _, slope := range slopes {
		if checkOnSlope(chairLayout, x, y, slope) {
			occupiedPlacesInSight++
		}
	}
	return occupiedPlacesInSight
}

func checkOnSlope(chairLayout [][]rune, x int, y int, slope Slope) bool {
	currentX, currentY := x, y
	for {
		currentX += slope.x
		currentY += slope.y
		if currentX < 0 || len(chairLayout[0]) <= currentX || currentY < 0 || len(chairLayout) <= currentY {
			return false
		}
		currentPlace := chairLayout[currentY][currentX]
		if currentPlace == CHAIR {
			return false
		}
		if currentPlace == OCCUPIED {
			return true
		}
	}
}
