package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Bus struct
type Bus struct {
	Pos int
	ID  int
}

// GetTimeTillNextBus calculates at what time the next bus will arive
func (b *Bus) GetTimeTillNextBus(time int) int {
	if time%b.ID == 0 {
		return 0
	}
	nextBus := math.Floor(float64(time/b.ID)) + 1
	nextBusTime := nextBus * float64(b.ID)
	return int(nextBusTime - float64(time))

}

var timestamp int = 0
var busses []Bus = []Bus{}

func main() {
	readFile()
	one()
	two()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile() {
	file, err := os.Open("./input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	timestampString := scanner.Text()
	timestamp = convertStringToInt(timestampString)
	scanner.Scan()
	busString := scanner.Text()
	busIDs := strings.Split(busString, ",")
	for i, IDstring := range busIDs {
		if IDstring != "x" {
			busID := convertStringToInt(IDstring)
			bus := Bus{ID: busID, Pos: i}
			busses = append(busses, bus)
		}
	}
	check(scanner.Err())
}

func convertStringToInt(str string) int {
	num, err := strconv.Atoi(str)
	check(err)
	return num
}

func one() {
	fastestTime := 10000
	fastestID := 10000
	for _, bus := range busses {
		time := bus.GetTimeTillNextBus(timestamp)
		if time < fastestTime {
			fastestTime = time
			fastestID = bus.ID
		}
	}
	fmt.Printf("Part one: %v\n", fastestID*fastestTime)
}

func two() {
	currentTime := 0
	step := 1
	for _, bus := range busses {
		for (currentTime+bus.Pos)%bus.ID != 0 {
			currentTime += step
		}
		step *= bus.ID
	}
	fmt.Printf("Part two: %v\n", currentTime)
}
