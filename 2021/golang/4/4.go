package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func main() {
	// one()
	two()
}

var (
	calls CallingOrder = CallingOrder{}
	boards []*Board
)

type MarkeableInt struct {
	Value int
	Marked bool
}

type Row struct {
	Values [5]*MarkeableInt
}

type Board struct {
	Rows [5]*Row
}

type CallingOrder struct {
	Values []int
}

func (num MarkeableInt) String() string {
	if (num.Value < 10) {
		if num.Marked {
			return fmt.Sprintf("--%d-", num.Value)
		}
		return fmt.Sprintf("  %d ", num.Value)
	}
	if num.Marked {
		return fmt.Sprintf("-%d-", num.Value)
	}
	return fmt.Sprintf(" %d ", num.Value)
}

func (b Board) String() string {
	output := ""
	for _, row := range b.Rows {
		for _, num := range row.Values {
			output += num.String()
		}
		output += "\n"
	}
	return output
}

func (b Board) MarkInteger(val int) bool {
	for _, row := range b.Rows {
		for _, num := range row.Values {
			if (num.Value == val) {
				num.Marked = true
				return true
			}
		}
	}
	return false
}

func (b Board) CheckWinCondition() bool {
	// tlToBrWin := true // TopLeft to BottomRight
	// blToTrWin := true // BottomLeft to TopRight
	for i, row := range b.Rows {
		if (row.CheckWinCondition()) {
			return true
		}
		columnWin := true
		for _, row := range b.Rows {
			if (!row.Values[i].Marked) {
				columnWin = false
			}
		}
		if (columnWin) {
			return true
		}
		// if (!b.Rows[i].Values[i].Marked) {
		// 	tlToBrWin = false
		// }
		// if (!b.Rows[4-i].Values[i].Marked) {
		// 	blToTrWin = false
		// }
	}
	// return tlToBrWin || blToTrWin
	return false
}

func (b Board) SumUnmarkedNumbers() int {
	total := 0
	for _, row := range b.Rows {
		for _, num := range row.Values {
			if (!num.Marked) {
				total += num.Value
			}
		}
	}
	return total
}

func (r Row) CheckWinCondition() bool {
	for _, markedNum := range r.Values {
		if (!markedNum.Marked) {
			return false;
		}
	}
	return true
}

func PrepareInput(file string) {
	input, _ := ioutil.ReadFile(file)
	lines := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	CreateCallingOrder(lines[0])
	for _, line := range lines[1:] {
		board := &Board{}
		rows := strings.Split(line, "\n")
		for i, row := range rows {
			numberStrings := strings.Split(row, " ")
			numbers := parseArrayOfStringsToArrayOfNumbers(numberStrings)
			row := InitializeRowFromIntegerArray(numbers)
			board.Rows[i] = row
		}
		boards = append(boards, board)
	}
}

func CreateCallingOrder(rawInput string) {
	numberStrings := strings.Split(rawInput, ",")
	numbers := parseArrayOfStringsToArrayOfNumbers(numberStrings)
	calls.Values = append(calls.Values, numbers...)
}

func parseArrayOfStringsToArrayOfNumbers(input []string) []int {
	returnList := []int{}
	input = sanatizeStringArray(input)
	for _, numString := range input {
		number, _ := strconv.Atoi(numString)
		returnList = append(returnList, number)
	}
	return returnList
}

func sanatizeStringArray(input []string) []string {
	returnList := []string{}
	for _, line := range input {
		line = strings.TrimSpace(line)
		if (len(line) > 0) {
			returnList = append(returnList, line)
		}
	}
	return returnList
}

func InitializeRowFromIntegerArray(numbers []int) *Row {
	r := &Row{}
	for i, num := range numbers {
		r.Values[i] = &MarkeableInt{
			Value: num,
			Marked: false,
		}
	}
	return r
}

func one() {
	PrepareInput("input.txt")
	for i, call := range calls.Values {
		fmt.Printf("==== %d - %d Called ====\n", i, call)
		for _, board := range boards {
			success := board.MarkInteger(call)
			if (success) {
				if (board.CheckWinCondition()) {
					fmt.Println("Winner Winner Chicken Dinner!")
					fmt.Println(board)
					fmt.Printf("Part one: %d", board.SumUnmarkedNumbers() * call)
					return;
				}
				
			}
			fmt.Println(board)
		}
	}
}

func two() {
	PrepareInput("input.txt")
	losingBoards := boards
	for i, call := range calls.Values {
		fmt.Printf("==== %d - %d Called ====\n", i, call)
		boards := losingBoards
		losingBoards = []*Board{}
		for _, board := range boards {
			board.MarkInteger(call)
			if (!board.CheckWinCondition()) {
				losingBoards = append(losingBoards, board)
			}
		fmt.Println(board)
		}
		if (len(losingBoards) == 0) {
			lastBoardToFinish := boards[len(boards)-1]
			fmt.Println("Loser Loser, *starts fortnite dancing*!")
			fmt.Println(lastBoardToFinish)
			fmt.Printf("Part two: %d\n", lastBoardToFinish.SumUnmarkedNumbers() * call)
			return
		}
	}
}