package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var (
	area       [][]int
	minRiskMap [][]int
	areaWidth  int
	areaHeight int
)

type Position struct {
	X int
	Y int
}

func (p Position) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

func main() {
	// one()
	two()
}

func prepareInput(file string) {
	input, _ := ioutil.ReadFile(file)
	inputStrings := strings.Split(strings.TrimSpace(string(input)), "\n")
	areaWidth = len(inputStrings[0])
	areaHeight = len(inputStrings)
	initAreas()
	for y, rowString := range inputStrings {
		for x, riskString := range rowString {
			risk, _ := strconv.Atoi(string(riskString))
			area[y][x] = risk
			minRiskMap[y][x] = 1000000000
		}
	}
	minRiskMap[0][0] = area[0][0]
}

func prepareEnlargedInput(file string) {
	input, _ := ioutil.ReadFile(file)
	inputStrings := strings.Split(strings.TrimSpace(string(input)), "\n")
	areaWidth = len(inputStrings[0]) * 5
	areaHeight = len(inputStrings) * 5
	initAreas()
	for i := 0; i < 5; i++ {
		for y, rowString := range inputStrings {
			for j := 0; j < 5; j++ {
				for x, riskString := range rowString {
					risk, _ := strconv.Atoi(string(riskString))
					risk = ((risk + i + j) % 9)
					if risk == 0 {
						risk = 1
					}
					area[y*(i+1)][x*(j+1)] = risk
					minRiskMap[y*(i+1)][x*(j+1)] = 1000000000
				}
			}
		}
	}
	minRiskMap[0][0] = area[0][0]
}

func initAreas() {
	area = make([][]int, areaWidth)
	minRiskMap = make([][]int, areaHeight)
	for i := range area {
		area[i] = make([]int, areaWidth)
		minRiskMap[i] = make([]int, areaHeight)
	}
}

func one() {
	prepareInput("input.txt")
	fillInMinRiskMap(0, 0)
	fmt.Printf("One: %d\n", minRiskMap[areaHeight-1][areaWidth-1]-area[0][0])
}

func two() {
	prepareEnlargedInput("example_input.txt")
	printArea()
}

func fillInMinRiskMap(fromX, fromY int) {
	adjecantPositions := getAdjecantPositions(fromX, fromY)
	fromRisk := minRiskMap[fromY][fromX]
	for _, adjecant := range adjecantPositions {
		adjecantRisk := area[adjecant.Y][adjecant.X]
		if (fromRisk + adjecantRisk) < minRiskMap[adjecant.Y][adjecant.X] {
			minRiskMap[adjecant.Y][adjecant.X] = fromRisk + adjecantRisk
			fillInMinRiskMap(adjecant.X, adjecant.Y)
		}
	}
}

func getAdjecantPositions(x, y int) []Position {
	adjecantPositions := []Position{}
	if x > 0 {
		adjecantPositions = append(adjecantPositions, Position{X: x - 1, Y: y})
	}
	if y > 0 {
		adjecantPositions = append(adjecantPositions, Position{X: x, Y: y - 1})
	}
	if x < areaWidth-1 {
		adjecantPositions = append(adjecantPositions, Position{X: x + 1, Y: y})
	}
	if y < areaHeight-1 {
		adjecantPositions = append(adjecantPositions, Position{X: x, Y: y + 1})
	}
	return adjecantPositions
}

func printArea() {
	output := make([][]string, areaHeight+1)
	for i := range output {
		output[i] = make([]string, areaWidth+1)
	}
	for y, row := range area {
		for x, risk := range row {
			output[y][x] = fmt.Sprint(risk)
		}
	}
	for _, line := range output {
		fmt.Println(line)
	}
}

func printRisk() {
	output := make([][]string, areaHeight+1)
	for i := range output {
		output[i] = make([]string, areaWidth+1)
	}
	for y, row := range minRiskMap {
		for x, risk := range row {
			output[y][x] = fmt.Sprint(risk)
		}
	}
	for _, line := range output {
		fmt.Println(strings.Join(line, "\t"))
	}
}
