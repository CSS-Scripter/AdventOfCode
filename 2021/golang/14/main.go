package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var (
	PolymerTemplate string
	Rules           map[string]byte
)

func main() {
	one()
	two()
}

func prepareInput(file string) {
	Rules = map[string]byte{}
	input, _ := ioutil.ReadFile(file)
	inputStrings := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	PolymerTemplate = inputStrings[0]
	for _, insertionRule := range strings.Split(inputStrings[1], "\n") {
		requirement := strings.Split(insertionRule, " -> ")[0]
		output := strings.Split(insertionRule, " -> ")[1][0]
		Rules[requirement] = output
	}
}

func one() {
	prepareInput("input.txt")
	step(10)
	mostCommonElement := 0
	leastCommonElement := 1000000
	for _, amount := range countElements() {
		if amount > mostCommonElement {
			mostCommonElement = amount
		}
		if amount < leastCommonElement {
			leastCommonElement = amount
		}
	}
	fmt.Printf("One %d\n", mostCommonElement-leastCommonElement)
}

func step(amount int) {
	for step := 0; step < amount; step++ {
		newPolymer := []byte{}
		for i := 1; i < len(PolymerTemplate); i++ {
			newPolymer = append(newPolymer, PolymerTemplate[i-1])
			newPolymer = append(newPolymer, Rules[PolymerTemplate[i-1:i+1]])
		}
		newPolymer = append(newPolymer, PolymerTemplate[len(PolymerTemplate)-1])
		PolymerTemplate = string(newPolymer)
	}
}

func countElements() map[byte]int {
	count := map[byte]int{}
	for _, el := range PolymerTemplate {
		count[byte(el)]++
	}
	return count
}

func two() {
	// prepareInput("input.txt")
	// step(40)
	// mostCommonElement := 0
	// leastCommonElementSet := false
	// leastCommonElement := 0
	// for _, amount := range countElements() {
	// 	if amount > mostCommonElement {
	// 		mostCommonElement = amount
	// 	}
	// 	if amount < leastCommonElement || !leastCommonElementSet {
	// 		leastCommonElement = amount
	// 		leastCommonElementSet = true
	// 	}
	// }
	// fmt.Printf("Two %d\n", mostCommonElement-leastCommonElement)
	fmt.Println("Two: fkn i dunno, shits crazy")
}
