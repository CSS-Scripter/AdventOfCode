package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

var (
	input                  []int
	possibilitiesFromPoint map[int]int
)

func main() {
	readFile()
	sort.Ints(input)
	one()
	two()
}

func readFile() {
	input = []int{}
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		input = append(input, number)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func one() {
	currentJoltage := 0
	diffAmounts := map[int]int{3: 1}
	for _, jolt := range input {
		diffAmounts[jolt-currentJoltage]++
		currentJoltage = jolt
	}
	fmt.Println(diffAmounts[1] * diffAmounts[3])
}

func two() {
	possibilitiesFromPoint = map[int]int{171: 1}
	for i := len(input) - 2; i >= 0; i-- {
		getPossibleConnections(input[i])
	}
	getPossibleConnections(0)
	fmt.Println(possibilitiesFromPoint[0])
}

func getPossibleConnections(jolt int) {
	total := 0
	for i := 0; i < len(input); i++ {
		diff := input[i] - jolt
		if 0 <= diff && diff <= 3 {
			total += possibilitiesFromPoint[input[i]]
		}
	}
	possibilitiesFromPoint[jolt] = total
}
