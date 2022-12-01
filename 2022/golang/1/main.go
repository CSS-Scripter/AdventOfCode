package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

var elfs []Elf

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
	sortElfs()
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
		elfs = append(elfs, elf)
	}
}

func sortElfs() {
	sort.Slice(elfs, func(i, j int) bool {
		return elfs[i].TotalCals > elfs[j].TotalCals
	})
}

func one() {
	fmt.Printf("1: %d\n", elfs[0].TotalCals)
}

func two() {
	total := elfs[0].TotalCals + elfs[1].TotalCals + elfs[2].TotalCals
	fmt.Printf("2: %d\n", total)
}
