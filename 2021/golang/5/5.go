package main

import (
	"io/ioutil"
	"strings"
	"fmt"
	"strconv"
)

var (
	vents []*Vent
	area [1000][1000]int = [1000][1000]int{}
	// area [10][10]int = [10][10]int{}
)

type Vent struct {
	fromX int
	fromY int
	toX int
	toY int
}

func (v Vent) String() string {
	return fmt.Sprintf("%d,%d -> %d,%d\n", v.fromX, v.fromY, v.toX, v.toY)
}

func main() {
	GetInput("input.txt")
	fmt.Println(vents)
	// getBiggestValuesFromVents()
	// drawAxialVents()
	drawAllVents()
	// printArea()
	countOverlaps()
}

func GetInput(file string) {
	input, _ := ioutil.ReadFile(file)
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	for _, line := range lines {
		values := strings.Split(line, " -> ")
		fromX, _ := strconv.Atoi(strings.Split(values[0], ",")[0])
		fromY, _ := strconv.Atoi(strings.Split(values[0], ",")[1])
	
		toX, _ := strconv.Atoi(strings.Split(values[1], ",")[0])
		toY, _ := strconv.Atoi(strings.Split(values[1], ",")[1])

		vents = append(vents, &Vent{
			fromX, fromY, toX, toY,
		})
	}
}

func drawAxialVents() {
	for _, vent := range vents {
		if (isVentHorizontal(vent)) {
			drawHorizontalVent(vent)
		}
		if (isVentVertical(vent)) {
			drawVerticalVent(vent)
		}
	}
}

func drawHorizontalVent(vent *Vent) {
	var (
		smallerX int
		biggerX int
	)
	if smallerX = vent.fromX; vent.fromX > vent.toX { smallerX = vent.toX }
	if biggerX = vent.fromX; vent.fromX < vent.toX { biggerX = vent.toX }
	for i := smallerX; i <= biggerX; i++ {
		area[i][vent.fromY] += 1;
	}
}

func drawVerticalVent(vent *Vent) {
	var (
		smallerY int
		biggerY int
	)
	if smallerY = vent.fromY; vent.fromY > vent.toY { smallerY = vent.toY }
	if biggerY = vent.fromY; vent.fromY < vent.toY { biggerY = vent.toY }
	for i := smallerY; i <= biggerY; i++ {
		area[vent.fromX][i] += 1;
	}
}

func isVentHorizontal(vent *Vent) bool {
	return vent.fromY == vent.toY && vent.fromX != vent.toX
}

func isVentVertical(vent *Vent) bool {
	return vent.fromX == vent.toX && vent.fromY != vent.toY
}

func countOverlaps() {
	overlaps := 0
	for _, verticalSlice := range area {
		for _, ventAmount := range verticalSlice {
			if (ventAmount > 1) {
				overlaps += 1
			}
		}
	}
	fmt.Printf("Amount of overlaps: %d", overlaps)
}

func drawAllVents() {
	for _, vent := range vents {
		if (isVentHorizontal(vent)) {
			drawHorizontalVent(vent)
		}
		if (isVentVertical(vent)) {
			drawVerticalVent(vent)
		}
		if (isVentDiagonal(vent)) {
			drawDiagonalVent(vent)
		}
	}
}

func isVentDiagonal(vent *Vent) bool {
	return !isVentHorizontal(vent) && !isVentVertical(vent)
}

func drawDiagonalVent(vent *Vent) {
	var (
		smallerX int
		biggerX int
		negativeX bool
		negativeY bool
		markX int
		markY int
	)
	if smallerX = vent.fromX; vent.fromX > vent.toX { smallerX = vent.toX }
	if biggerX = vent.fromX; vent.fromX < vent.toX { biggerX = vent.toX }

	negativeY = vent.fromY > vent.toY;
	negativeX = vent.fromX > vent.toX;

	delta := biggerX - smallerX;

	for i := 0; i <= delta; i++ {
		if (negativeX) {markX = vent.fromX - i} else {markX = vent.fromX + i}
		if (negativeY) {markY = vent.fromY - i} else {markY = vent.fromY + i}
		area[markX][markY] += 1;
	}
}

// Only used to get the area size
func getBiggestValuesFromVents() {
	biggestX := 0
	biggestY := 0

	for _, vent := range vents {
		if (vent.fromX > biggestX) { biggestX = vent.fromX }
		if (vent.fromY > biggestY) { biggestY = vent.fromY }
		if (vent.toX > biggestX) { biggestX = vent.toX }
		if (vent.toY > biggestY) { biggestY = vent.toY }
	}
	fmt.Printf("Biggest X %d\n", biggestX)
	fmt.Printf("Biggest Y %d\n", biggestY)
}

func printArea() {
	lines := [len(area)]string{}
	for _, verticalSlice := range area {
		for i, ventAmount := range verticalSlice {
			lines[i] += fmt.Sprintf("%d ", ventAmount)
		}
	}
	for _, line := range lines {
		fmt.Printf("%s\n", line)
	}
}

