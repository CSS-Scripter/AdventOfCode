package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

var (
	traversedLocations map[int]map[int]int
	paths              []Path
)

type Movement struct {
	Dir    byte
	Amount int
}

func (m Movement) String() string {
	return fmt.Sprintf("%s%d", string(m.Dir), m.Amount)
}

func (m Movement) MoveFrom(x, y int) (int, int) {
	switch m.Dir {
	case 'U':
		if traversedLocations[x] == nil {
			traversedLocations[x] = map[int]int{}
		}
		for i := 1; i <= m.Amount; i++ {
			traversedLocations[x][y+i] += 1
		}
		return x, y + m.Amount

	case 'D':
		if traversedLocations[x] == nil {
			traversedLocations[x] = map[int]int{}
		}
		for i := 1; i <= m.Amount; i++ {
			traversedLocations[x][y-i] += 1
		}
		return x, y - m.Amount

	case 'L':
		for i := 1; i <= m.Amount; i++ {
			if traversedLocations[x-i] == nil {
				traversedLocations[x-i] = map[int]int{}
			}
			traversedLocations[x-i][y] += 1
		}
		return x - m.Amount, y

	case 'R':
		for i := 1; i <= m.Amount; i++ {
			if traversedLocations[x+i] == nil {
				traversedLocations[x+i] = map[int]int{}
			}
			traversedLocations[x+i][y] += 1
		}
		return x + m.Amount, y
	}
	panic("Illegal state reached!")
}

type Path struct {
	Moves []Movement
}

func main() {
	one()
	two()
}

func one() {
	prepareInput("input.txt")
	for _, path := range paths {
		x := 0
		y := 0
		for _, movement := range path.Moves {
			x, y = movement.MoveFrom(x, y)
		}
	}
	smallestDistance := 10000
	for x, yMap := range traversedLocations {
		for y, traversedAmount := range yMap {
			manhattanDistance := int(math.Abs(float64(x)) + math.Abs(float64(y)))
			if traversedAmount > 1 && manhattanDistance < smallestDistance {
				smallestDistance = manhattanDistance
			}
		}
	}
	fmt.Printf("One: %d\n", smallestDistance)
}

func two() {
	fmt.Printf("Two: %d\n", -1)
}

func prepareInput(file string) {
	traversedLocations = make(map[int]map[int]int)
	paths = []Path{}
	input, _ := ioutil.ReadFile(file)
	inputStrings := strings.Split(strings.TrimSpace(string(input)), "\r\n")
	for _, pathString := range inputStrings {
		path := Path{}
		for _, movString := range strings.Split(pathString, ",") {
			dir := []byte(movString)[0]
			amount, _ := strconv.Atoi(string([]byte(movString)[1:]))
			movement := Movement{Dir: dir, Amount: amount}
			path.Moves = append(path.Moves, movement)
		}
		paths = append(paths, path)
	}
}
