package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var (
	area   Map
	folds  []Fold
	width  int
	height int
)

type Map []Coord

func (m Map) String() string {
	stringMap := [][]byte{}
	for y := 0; y <= height; y++ {
		line := []byte{}
		for x := 0; x <= width; x++ {
			line = append(line, '.')
		}
		stringMap = append(stringMap, line)
	}
	for _, pos := range m {
		stringMap[pos.Y][pos.X] = '#'
	}
	returnString := ""
	for _, line := range stringMap {
		returnString += (string(line) + "\n")
	}
	return returnString
}

func (m Map) getMaxY() int {
	biggestY := 0
	for _, pos := range m {
		if pos.Y > biggestY {
			biggestY = pos.Y
		}
	}
	return biggestY
}

func (m Map) getMaxX() int {
	biggestX := 0
	for _, pos := range m {
		if pos.X > biggestX {
			biggestX = pos.X
		}
	}
	return biggestX
}

func (m Map) Fold(f Fold) Map {
	newMap := Map{}
	for _, pos := range m {
		newMap = append(newMap, pos.GetMirroredCoord(f))
	}
	updateWidth(f)
	updateHeight(f)
	return newMap
}

func updateWidth(f Fold) {
	if f.Dir == 'y' {
		return
	}
	if f.Pos%2 == 1 {
		width = f.Pos - 1
	} else {
		width = f.Pos
	}
}

func updateHeight(f Fold) {
	if f.Dir == 'x' {
		return
	}
	if f.Pos%2 == 1 {
		height = f.Pos - 1
	} else {
		height = f.Pos
	}
}

type Coord struct {
	X int
	Y int
}

type Fold struct {
	Dir byte
	Pos int
}

func (f Fold) String() string {
	return fmt.Sprintf("%s %d", string(f.Dir), f.Pos)
}

func (c Coord) GetMirroredCoord(fold Fold) Coord {
	switch fold.Dir {
	case 'x':
		return c.mirrorVertical(fold.Pos)
	case 'y':
		return c.mirrorHorizontal(fold.Pos)
	}
	panic(fmt.Sprintf("Invalid direction byte %#v", fold))
}

func (c Coord) mirrorVertical(pos int) Coord {
	newX := c.X
	if c.X > pos {
		newX = pos - (c.X - pos)
	}
	return Coord{X: newX, Y: c.Y}
}

func (c Coord) mirrorHorizontal(pos int) Coord {
	newY := c.Y
	if c.Y > pos {
		newY = pos - (c.Y - pos)
	}
	return Coord{X: c.X, Y: newY}
}

func main() {
	one()
	two()
}

func one() {
	prepareInput("input.txt")
	area = area.Fold(folds[0])
	areaString := area.String()
	dotCounter := 0
	for _, char := range areaString {
		if char == '#' {
			dotCounter += 1
		}
	}
	fmt.Printf("One: %d\n", dotCounter)
}

func two() {
	prepareInput("input.txt")
	for _, f := range folds {
		area = area.Fold(f)
	}
	areaString := area.String()
	fmt.Println("Two: FPEKBEJL")
	fmt.Println(areaString)
}

func prepareInput(file string) {
	input, _ := ioutil.ReadFile(file)
	inputStrings := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	for _, coordsString := range strings.Split(inputStrings[0], "\n") {
		coords := strings.Split(coordsString, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		area = append(area, Coord{X: x, Y: y})
	}
	for _, foldString := range strings.Split(inputStrings[1], "\n") {
		foldsData := strings.Split(strings.TrimPrefix(foldString, "fold along "), "=")
		dir := foldsData[0][0]
		pos, _ := strconv.Atoi(foldsData[1])
		folds = append(folds, Fold{Dir: dir, Pos: pos})
	}
	width = area.getMaxX()
	height = area.getMaxY()
}

func getSmallestXFold() int {
	smallestXFold := 10000
	for _, f := range folds {
		if f.Dir == 'x' && f.Pos < smallestXFold {
			smallestXFold = f.Pos
		}
	}
	return smallestXFold
}

func getSmallestYFold() int {
	smallestYFold := 10000
	for _, f := range folds {
		if f.Dir == 'y' && f.Pos < smallestYFold {
			smallestYFold = f.Pos
		}
	}
	return smallestYFold
}
