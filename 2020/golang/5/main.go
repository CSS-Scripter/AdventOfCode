package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

const (
	maxRows     = 127
	rowSteps    = 7
	maxColumns  = 7
	columnSteps = 3
)

// Seat contains information
type Seat struct {
	RowSequence    string
	ColumnSequence string
	Row            int
	Column         int
	ID             int
}

// CalculateRow calculates the row based on the set RowSequence of the seat
func (s *Seat) CalculateRow() {
	currentValue := 0
	currentHalfMin := 0
	currentHalfMax := maxRows
	for step := 1; step <= rowSteps; step++ {
		character := rune(s.RowSequence[step-1])
		if character == 'F' {
			currentHalfMax -= (currentHalfMax-currentHalfMin)/2 + 1
		}
		if character == 'B' {
			currentHalfMin += (currentHalfMax-currentHalfMin)/2 + 1
			currentValue = currentHalfMin
		}
	}
	s.Row = currentValue
}

// CalculateColumn calculates the column based on the set ColumnSequence of the seat
func (s *Seat) CalculateColumn() {
	currentValue := 0
	currentHalfMin := 0
	currentHalfMax := maxColumns
	for step := 1; step <= columnSteps; step++ {
		character := rune(s.ColumnSequence[step-1])
		if character == 'L' {
			currentHalfMax -= (currentHalfMax-currentHalfMin)/2 + 1
		}
		if character == 'R' {
			currentHalfMin += (currentHalfMax-currentHalfMin)/2 + 1
			currentValue = currentHalfMin
		}
	}
	s.Column = currentValue
}

// CalculateID calculates the ID based on the set row and column of the seat
func (s *Seat) CalculateID() {
	s.ID = (s.Row * 8) + s.Column
}

func main() {
	seats := readFile()
	first(seats)
	second(seats)
}

func first(seats []Seat) {
	maxID := 0
	for _, seat := range seats {
		if seat.ID > maxID {
			maxID = seat.ID
		}
	}
	fmt.Printf("The highest ID found is %v\n", maxID)
}

func second(seats []Seat) {
	foundIDs := []int{}
	for _, seat := range seats {
		foundIDs = append(foundIDs, seat.ID)
	}
	sort.Ints(foundIDs)
	previousID := foundIDs[0]
	for _, ID := range foundIDs {
		if ID-previousID == 2 {
			fmt.Printf("Found missing ID: %v\n", previousID+1)
		}
		previousID = ID
	}
}

func readFile() []Seat {
	returnValue := []Seat{}

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		seat := Seat{}
		line := scanner.Text()
		seat.RowSequence = line[0:7]
		seat.ColumnSequence = line[7:]
		seat.CalculateRow()
		seat.CalculateColumn()
		seat.CalculateID()
		returnValue = append(returnValue, seat)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return returnValue
}
