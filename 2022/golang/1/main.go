package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

var elfes []Elf

type Elf struct {
	Calories  []int
	TotalCals int
}

func (e *Elf) CalculateTotalCalories() int {
	total := 0
	for _, cal := range e.Calories {
		total += cal
	}
	e.TotalCals = total
	return total
}

func main() {
	prepareInput("input.txt")
	one()
	two()
}

func prepareInput(file string) {
	input, _ := ioutil.ReadFile(file)
	lines := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	for _, line := range lines {
		var elf Elf
		for _, cal := range strings.Split(strings.TrimSpace(line), "\n") {
			calInt, _ := strconv.Atoi(cal)
			elf.Calories = append(elf.Calories, calInt)
		}
		elf.CalculateTotalCalories()
		elfes = append(elfes, elf)
	}
}

func one() {
	max := 0
	for _, elf := range elfes {
		if elf.TotalCals > max {
			max = elf.TotalCals
		}
	}
	fmt.Printf("1: %d\n", max)
}

func two() {
	sort.Slice(elfes, func(i, j int) bool {
		return elfes[i].TotalCals > elfes[j].TotalCals
	})
	total := elfes[0].TotalCals + elfes[1].TotalCals + elfes[2].TotalCals
	fmt.Printf("2: %d\n", total)
}
