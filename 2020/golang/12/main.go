package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// These constants are defined to remove magic numbers
// EAST is 0 because it is the beginning direction we are facing
// Turning right 90 degrees will increase direction by 1, turning left by -1
const (
	EAST  = 0
	SOUTH = 1
	WEST  = 2
	NORTH = 3
)

var (
	input    []Movement = []Movement{}
	curDir   int        = 0
	curNorth int        = 0
	curEast  int        = 0
	waypoint WayPoint   = WayPoint{East: 10, North: 1}
)

// UnknownMovementError is thrown when there is an unknown movement found
type UnknownMovementError struct {
	Move Movement
}

func (err *UnknownMovementError) Error() string {
	return fmt.Sprintf("Unknown movement occured: %v", err.Move.String())
}

// Movement is used to define a single movement, which would be a single line from the given input
type Movement struct {
	Direction rune
	Amount    int
}

// Execute executes the current
func (move Movement) Execute() {
	switch move.Direction {
	case 'N':
		curNorth += move.Amount
		break
	case 'E':
		curEast += move.Amount
		break
	case 'S':
		curNorth -= move.Amount
		break
	case 'W':
		curEast -= move.Amount
		break
	case 'R':
		curDir += move.Amount / 90
		curDir %= 4
		break
	case 'L':
		curDir -= move.Amount / 90
		if curDir < 0 {
			curDir += 4
		}
		break
	case 'F':
		switch curDir {
		case NORTH:
			curNorth += move.Amount
			break
		case EAST:
			curEast += move.Amount
			break
		case SOUTH:
			curNorth -= move.Amount
			break
		case WEST:
			curEast -= move.Amount
			break
		default:
			fmt.Printf("Unknown direction occured: %v\n", curDir)
		}
		break
	default:
		panic(UnknownMovementError{Move: move})
	}
}

// ExecuteTwo executes movement conform the second challenge.
func (move Movement) ExecuteTwo() {
	switch move.Direction {
	case 'N':
		waypoint.North += move.Amount
		break
	case 'E':
		waypoint.East += move.Amount
		break
	case 'S':
		waypoint.North -= move.Amount
		break
	case 'W':
		waypoint.East -= move.Amount
		break
	case 'R':
		waypoint.RotateRight(move.Amount)
		break
	case 'L':
		waypoint.RotateLeft(move.Amount)
		break
	case 'F':
		curNorth += (waypoint.North * move.Amount)
		curEast += (waypoint.East * move.Amount)
		break
	default:
		panic(UnknownMovementError{Move: move})
	}
}

func (move *Movement) String() string {
	return fmt.Sprintf("Movement<Direction: %v(%v), Amount: %v>", move.Direction, string(move.Direction), move.Amount)
}

// WayPoint defines the direction the ship moves towards
type WayPoint struct {
	North int
	East  int
}

// RotateLeft rotates the waypoint for a certain amount of degrees
func (wp *WayPoint) RotateLeft(deg int) {
	wp.RotateRight(-deg)
}

// RotateRight rotates the waypoint for a certain amount of degrees
func (wp *WayPoint) RotateRight(deg int) {
	amount := deg / 90
	amount %= 4
	if amount < 0 {
		amount += 4
	}
	for i := 0; i < amount; i++ {
		newNorth := -wp.East
		wp.East = wp.North
		wp.North = newNorth
	}
}

func main() {
	readFile()
	one()
	two()
}

func readFile() {
	file, err := os.Open("./input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		direction := rune(line[0])
		amount, err := strconv.Atoi(line[1:])
		check(err)
		movement := Movement{
			Direction: direction,
			Amount:    amount,
		}
		input = append(input, movement)
	}
	check(scanner.Err())
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func one() {
	for _, movement := range input {
		movement.Execute()
	}
	printAbsoluteNorthEastValues()
}

func printAbsoluteNorthEastValues() {
	north := curNorth
	east := curEast
	if curNorth >= 0 {
		fmt.Printf("North by %v\n", north)
	} else {
		north = -north
		fmt.Printf("South by %v\n", north)
	}
	if curEast >= 0 {
		fmt.Printf("East by %v\n", east)
	} else {
		east = -east
		fmt.Printf("West by %v\n", east)
	}
	fmt.Printf("Manhattan Distance: %v\n", east+north)
}

func two() {
	curDir = 0
	curEast = 0
	curNorth = 0
	for _, movement := range input {
		movement.ExecuteTwo()
	}
	printAbsoluteNorthEastValues()
}
