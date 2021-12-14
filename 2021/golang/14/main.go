package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var (
	PolymerTemplate map[string]int
	Rules           map[string]string
	StartingElement byte
	EndingElement   byte
)

func main() {
	one()
	two()
}

func prepareInput(file string) {
	PolymerTemplate = map[string]int{}
	Rules = map[string]string{}
	input, _ := ioutil.ReadFile(file)
	inputStrings := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	StartingElement = inputStrings[0][0]
	EndingElement = inputStrings[0][len(inputStrings[0])-1]
	for i := 1; i < len(inputStrings[0]); i++ {
		PolymerTemplate[inputStrings[0][i-1:i+1]] += 1
	}
	for _, insertionRule := range strings.Split(inputStrings[1], "\n") {
		requirement := strings.Split(insertionRule, " -> ")[0]
		output := strings.Split(insertionRule, " -> ")[1]
		Rules[requirement] = output
	}
}

func one() {
	prepareInput("example_input.txt")
	step(10)
	mostCommonElement := 0
	leastCommonElement := 0
	leastCommonElementSet := false
	for _, amount := range countElements() {
		if amount > mostCommonElement {
			mostCommonElement = amount
		}
		if amount < leastCommonElement || !leastCommonElementSet {
			leastCommonElement = amount
			leastCommonElementSet = true
		}
	}
	fmt.Printf("One %d\n", mostCommonElement-leastCommonElement)
}

func step(amount int) {
	for step := 0; step < amount; step++ {
		newPolymer := map[string]int{}
		for poly, amount := range PolymerTemplate {
			addition := Rules[poly]
			newPolyOne := string(poly[0]) + addition
			newPolyTwo := addition + string(poly[1])
			newPolymer[newPolyOne] += amount
			newPolymer[newPolyTwo] += amount
		}
		PolymerTemplate = newPolymer
	}
}

func countElements() map[byte]int {
	count := map[byte]int{}
	for poly, amount := range PolymerTemplate {
		count[poly[0]] += amount
		count[poly[1]] += amount
	}
	count[StartingElement] += 1
	count[EndingElement] += 1
	for el, elCount := range count {
		count[el] = elCount / 2
	}
	return count
}

func two() {
	prepareInput("input.txt")
	step(40)
	mostCommonElement := 0
	leastCommonElementSet := false
	leastCommonElement := 0
	for _, amount := range countElements() {
		if amount > mostCommonElement {
			mostCommonElement = amount
		}
		if amount < leastCommonElement || !leastCommonElementSet {
			leastCommonElement = amount
			leastCommonElementSet = true
		}
	}
	fmt.Printf("Two %d\n", mostCommonElement-leastCommonElement)
}
