package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

var (
	heightMap [][]int
)

type HeightPoint struct {
	height   int
	lowpoint bool
}

type Position struct {
	x int
	y int
}

func (p Position) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}

func main() {
	two()
}

func one() {
	prepareInput("input.txt")
	lowPoints := findLowPoints()
	sumRiskLevels(lowPoints)
}

func two() {
	prepareInput("input.txt")
	lowPoints := findLowPoints()
	calculateBasinSizes(lowPoints)
}

func prepareInput(file string) {
	input, _ := ioutil.ReadFile(file)
	inputStrings := strings.Split(strings.TrimSpace(string(input)), "\n")
	for _, inputLine := range inputStrings {
		heightLine := []int{}
		for _, char := range strings.TrimSpace(inputLine) {
			height, _ := strconv.Atoi(string(char))
			heightLine = append(heightLine, height)
		}
		heightMap = append(heightMap, heightLine)
	}
}

func findLowPoints() []Position {
	lowPoints := []Position{}
	for y, heightLine := range heightMap {
		for x := range heightLine {
			if isAdjecantPositionLower(x, y) {
				continue
			}
			lowPoints = append(lowPoints, Position{x, y})
		}
	}
	return lowPoints
}

func isAdjecantPositionLower(x int, y int) bool {
	height := heightMap[y][x]
	heightLine := heightMap[y]
	if y > 0 && heightMap[y-1][x] <= height {
		return true
	}
	if y < (len(heightMap)-1) && heightMap[y+1][x] <= height {
		return true
	}
	if x > 0 && heightMap[y][x-1] <= height {
		return true
	}
	if x < (len(heightLine)-1) && heightLine[x+1] <= height {
		return true
	}
	return false
}

func sumRiskLevels(lowPoints []Position) {
	riskLevel := 0
	for _, lowPoint := range lowPoints {
		height := heightMap[lowPoint.y][lowPoint.x]
		riskLevel += (height + 1)
		fmt.Printf("(%d, %d) = %d\n", lowPoint.x, lowPoint.y, height)
	}
	fmt.Printf("Risk level: %d\n", riskLevel)
}

func printHeightMap() {
	for _, heightLine := range heightMap {
		for _, height := range heightLine {
			fmt.Printf("%d", height)
		}
		fmt.Println()
	}
}

type Basin struct {
	positions []Position
}

func (b Basin) String() string {
	returnString := "["
	for _, pos := range b.positions {
		returnString += pos.String()
		returnString += ", "
	}
	returnString = returnString[:len(returnString)-2] + "]\n"
	return returnString
}

func (b Basin) containsPosition(position Position) bool {
	for _, basinPos := range b.positions {
		if position.x == basinPos.x && position.y == basinPos.y {
			return true
		}
	}
	return false
}

func calculateBasinSizes(lowPoints []Position) {
	var basins []*Basin
	for _, lowPoint := range lowPoints {
		if doesBasinCollectionContainLowPoint(basins, lowPoint) {
			continue
		}
		basin := &Basin{positions: []Position{}}
		calculateBasin(basin, lowPoint)
		basins = append(basins, basin)
	}
	basins = sortBasins(basins)
	fmt.Println(basins)
	total := 1
	for _, largestBasin := range basins[:3] {
		amountOfPositions := len(largestBasin.positions)
		fmt.Println(amountOfPositions)
		total *= amountOfPositions
	}
	fmt.Println(total)
}

func doesBasinCollectionContainLowPoint(basins []*Basin, lowPoint Position) bool {
	for _, basin := range basins {
		if basin.containsPosition(lowPoint) {
			return true
		}
	}
	return false
}

func calculateBasin(basin *Basin, p Position) {
	basin.positions = append(basin.positions, p)
	heightLine := heightMap[p.y]
	if p.y > 0 && heightMap[p.y-1][p.x] != 9 && !basin.containsPosition(Position{x: p.x, y: p.y - 1}) {
		calculateBasin(basin, Position{x: p.x, y: p.y - 1})
	}
	if p.y < (len(heightMap)-1) && heightMap[p.y+1][p.x] != 9 && !basin.containsPosition(Position{x: p.x, y: p.y + 1}) {
		calculateBasin(basin, Position{x: p.x, y: p.y + 1})
	}
	if p.x > 0 && heightMap[p.y][p.x-1] != 9 && !basin.containsPosition(Position{x: p.x - 1, y: p.y}) {
		calculateBasin(basin, Position{x: p.x - 1, y: p.y})
	}
	if p.x < (len(heightLine)-1) && heightLine[p.x+1] != 9 && !basin.containsPosition(Position{x: p.x + 1, y: p.y}) {
		calculateBasin(basin, Position{x: p.x + 1, y: p.y})
	}
}

func sortBasins(basins []*Basin) []*Basin {
	sort.SliceStable(basins, func(i, j int) bool {
		return len(basins[i].positions) > len(basins[j].positions)
	})
	return basins
}
