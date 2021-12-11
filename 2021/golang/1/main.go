package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var (
	heights []int
)

func main() {
	prepareInput("input.txt")
	one()
	two()
}

func one() {
	increments := 0
	for i, height := range heights[1:] {
		if height > heights[i] {
			increments += 1
		}
	}
	fmt.Printf("One: %d\n", increments)
}

func two() {
	increments := 0
	previousSum := heights[0] + heights[1] + heights[2]
	for i := range heights[3:] {
		newSum := heights[i+1] + heights[i+2] + heights[i+3]
		if newSum > previousSum {
			increments += 1
		}
		previousSum = newSum
	}
	fmt.Printf("Two: %d\n", increments)
}

func prepareInput(file string) {
	input, _ := ioutil.ReadFile(file)
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	for _, line := range lines {
		height, _ := strconv.Atoi(line)
		heights = append(heights, height)
	}
}
