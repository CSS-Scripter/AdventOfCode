package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
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
	start := time.Now()

	one()
	two()

	elapsed := time.Since(start)
	fmt.Printf("Duration: %s", elapsed)
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
			minRiskMap[y][x] = -1
			if x > 0 {
				minRiskMap[y][x] = minRiskMap[y][x-1] + risk
			} else if y > 0 {
				minRiskMap[y][x] = minRiskMap[y-1][x] + risk
			} else {
				minRiskMap[y][x] = risk
			}
		}
	}
	minRiskMap[0][0] = area[0][0]
}

func prepareEnlargedInput(file string) {
	input, _ := ioutil.ReadFile(file)
	inputStrings := strings.Split(strings.TrimSpace(string(input)), "\n")
	sectionWidth := len(inputStrings[0])
	sectionHeight := len(inputStrings)
	areaWidth = len(inputStrings[0]) * 5
	areaHeight = len(inputStrings) * 5
	initAreas()
	for i := 0; i < 5; i++ {
		for y, rowString := range inputStrings {
			for j := 0; j < 5; j++ {
				for x, riskString := range rowString {
					risk, _ := strconv.Atoi(string(riskString))
					risk += i + j
					if risk >= 10 {
						risk = risk - 9
					}
					totalX := x + (sectionWidth * j)
					totalY := y + (sectionHeight * i)
					area[totalY][totalX] = risk
					if totalX > 0 {
						minRiskMap[totalY][totalX] = minRiskMap[totalY][totalX-1] + risk
					} else if totalY > 0 {
						minRiskMap[totalY][totalX] = minRiskMap[totalY-1][totalX] + risk
					} else {
						minRiskMap[totalY][totalX] = risk
					}
				}
			}
		}
	}
}

func initAreas() {
	area = make([][]int, areaHeight)
	minRiskMap = make([][]int, areaHeight)
	for i := range area {
		area[i] = make([]int, areaWidth)
		minRiskMap[i] = make([]int, areaWidth)
	}
}

func one() {
	prepareInput("input.txt")
	fillInMinMapInverse()
	fmt.Printf("One: %d\n", minRiskMap[areaHeight-1][areaWidth-1]-area[0][0])
}

func two() {
	prepareEnlargedInput("input.txt")
	fillInMinMapInverse()
	fmt.Printf("Two: %d\n", minRiskMap[areaHeight-1][areaWidth-1]-area[0][0])
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
	output := make([][]string, areaHeight)
	for i := range output {
		output[i] = make([]string, areaWidth)
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
	output := make([][]string, areaHeight)
	for i := range output {
		output[i] = make([]string, areaWidth)
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

func fillInMinMapInverse() {
	changes := 0

	for y, row := range minRiskMap {
		for x, minRisk := range row {
			for _, ad := range getAdjecantPositions(x, y) {
				if (minRiskMap[ad.Y][ad.X] + area[y][x]) < minRisk {
					minRiskMap[y][x] = minRiskMap[ad.Y][ad.X] + area[y][x]
					changes += 1
				}
			}
		}
	}
	if changes > 0 {
		fillInMinMapInverse()
	}
}
