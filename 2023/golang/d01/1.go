package d01

import (
	"aoc2023/src/util"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
)

type NumberPosition struct {
	Position int
	Value    string
}

func Main() {
	input, err := util.ReadInput(1)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(input), "\n")
	one(lines)
	two(lines)
}

func one(input []string) {
	sum := 0
	for i, line := range input {
		numPositions := getNumberPositionsOfString(line)
		sort.Slice(numPositions[:], func(i, j int) bool {
			return numPositions[i].Position < numPositions[j].Position
		})
		first := numPositions[0]
		last := numPositions[len(numPositions)-1]
		value, err := strconv.Atoi(fmt.Sprintf("%s%s", first.Value, last.Value))
		if err != nil {
			log.Error(fmt.Sprintf("failed to parse value at line %d with first value %s and last value %s", i, first.Value, last.Value))
			panic(err)
		}
		sum += value
	}
	log.Info(fmt.Sprintf("part 1 solution: %d", sum))
}

func two(input []string) {
	sum := 0
	for i, line := range input {
		numPositions := getLiteralNumberPositionsOfString(line)
		sort.Slice(numPositions[:], func(i, j int) bool {
			return numPositions[i].Position < numPositions[j].Position
		})
		first := numPositions[0]
		last := numPositions[len(numPositions)-1]
		value, err := strconv.Atoi(fmt.Sprintf("%s%s", first.Value, last.Value))

		if err != nil {
			log.Error(fmt.Sprintf("failed to parse value at line %d with first value %s and last value %s", i, first.Value, last.Value))
			panic(err)
		}
		sum += value
	}
	log.Info(fmt.Sprintf("part 2 solution: %d", sum))
}

func getLiteralNumberPositionsOfString(input string) []NumberPosition {
	numPositions := getNumberPositionsOfString(input)
	literalNumbers := strings.Split("zero,one,two,three,four,five,six,seven,eight,nine", ",")
	for i, litNum := range literalNumbers {
		positions := getIndexesOfSubstring(input, litNum)
		for _, pos := range positions {
			numPositions = append(numPositions, NumberPosition{
				Value:    fmt.Sprintf("%d", i),
				Position: pos,
			})
		}
	}
	return numPositions
}

func getNumberPositionsOfString(input string) []NumberPosition {
	numPositions := []NumberPosition{}
	numerals := strings.Split("1234567890", "")
	for _, num := range numerals {
		positions := getIndexesOfSubstring(input, num)
		for _, pos := range positions {
			numPositions = append(numPositions, NumberPosition{
				Value:    num,
				Position: pos,
			})
		}
	}
	return numPositions
}

func getIndexesOfSubstring(input, substring string) []int {
	if !strings.Contains(input, substring) {
		return []int{}
	}

	index := strings.Index(input, substring)
	newInput := input[index+1:]
	nextValues := getIndexesOfSubstring(newInput, substring)
	for i := range nextValues {
		nextValues[i] += index + 1
	}
	return append(nextValues, index)
}
