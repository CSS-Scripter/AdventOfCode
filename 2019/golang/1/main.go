package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

var (
	masses []Mass
)

type Mass int

func (m Mass) calcFuel() int {
	return Max(0, int(math.Floor(float64(m)/3))-2)
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {
	one()
	two()
}

func one() {
	prepareInput("input.txt")
	totalFuel := 0
	for _, mass := range masses {
		totalFuel += mass.calcFuel()
	}
	fmt.Printf("One: %d\n", totalFuel)
}

func two() {
	prepareInput("input.txt")
	totalFuel := 0
	for {
		newMasses := []Mass{}
		fuelCount := 0
		for _, mass := range masses {
			fuel := mass.calcFuel()
			fuelCount += fuel
			newMasses = append(newMasses, Mass(fuel))
		}
		totalFuel += fuelCount
		if fuelCount == 0 {
			break
		}
		masses = newMasses
	}
	fmt.Printf("Two: %d\n", totalFuel)
}

func prepareInput(file string) {
	masses = []Mass{}
	input, _ := ioutil.ReadFile(file)
	inputStrings := strings.Split(strings.TrimSpace(string(input)), "\r\n")
	for _, inputLine := range inputStrings {
		mass, _ := strconv.Atoi(inputLine)
		masses = append(masses, Mass(mass))
	}
}
