package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	input := readFile()
	encounteredTrees := getEncounteredTrees(input, 1, 1)
	encounteredTrees *= getEncounteredTrees(input, 3, 1)
	encounteredTrees *= getEncounteredTrees(input, 5, 1)
	encounteredTrees *= getEncounteredTrees(input, 7, 1)
	encounteredTrees *= getEncounteredTrees(input, 1, 2)
	fmt.Printf("Encountered Trees: %d\n", encounteredTrees)

}

func readFile() [][]rune {
	returnValue := [][]rune{}

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := []rune{}
		for _, character := range scanner.Text() {
			row = append(row, character)
		}
		returnValue = append(returnValue, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return returnValue
}

func getEncounteredTrees(chart [][]rune, xSlope int, ySlope int) int {
	treesEncountered := 0
	steps := 0
	for y := ySlope; y < len(chart); y += ySlope {
		steps++
		x := steps * xSlope % len(chart[y])
		if chart[y][x] == '#' {
			treesEncountered++
		}
	}
	return treesEncountered
}
