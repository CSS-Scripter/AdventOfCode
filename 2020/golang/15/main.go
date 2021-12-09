package main

import "fmt"

var (
	input          = []int{13, 16, 0, 12, 15, 1}
	turn           = 0
	lastTimeSpoken = map[int]int{}
	lastNumber     = 0
)

func main() {
	one(2020)
	one(30000000) // Two
}

func one(rounds int) {
	lastTimeSpoken = map[int]int{}
	lastNumber = 0
	for i, in := range input {
		lastTimeSpoken[in] = i + 1
	}
	for turn = len(input) + 1; turn < rounds; turn++ {
		if lastTimeSpoken[lastNumber] == 0 {
			lastTimeSpoken[lastNumber] = turn
			lastNumber = 0
			continue
		}
		number := lastNumber
		lastNumber = turn - lastTimeSpoken[number]
		lastTimeSpoken[number] = turn
	}
	fmt.Println(lastNumber)
}
