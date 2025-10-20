package d02

import (
	"aoc2023/src/util"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
)

const (
	Red   string = "red"
	Blue  string = "blue"
	Green string = "green"
)

type Game struct {
	ID        int
	DiceGrabs []DiceGrab
}

type DiceGrab []Dices

type Dices struct {
	Count int
	Color string
}

func Main() {
	input, err := util.ReadInput(2)
	if err != nil {
		log.Error("failed to load input")
		panic(err)
	}

	lines := strings.Split(string(input), "\n")
	games := parseInput(lines)
	one(games)
	two(games)
}

func parseInput(lines []string) []Game {
	games := []Game{}
	for _, line := range lines {
		games = append(games, parseLine(line))
	}
	return games
}

func parseLine(line string) Game {
	segments := strings.Split(line, ": ")
	if len(segments) != 2 {
		log.Error(fmt.Sprintf("line resulted in too many segments: %s", line))
		panic(errors.New("panic!"))
	}

	return Game{
		ID:        getGameIdFromSegment(segments[0]),
		DiceGrabs: parseDiceGrabs(segments[1]),
	}
}

func parseDiceGrabs(segment string) []DiceGrab {
	diceGrabs := []DiceGrab{}
	segments := strings.Split(segment, "; ")
	for _, seg := range segments {
		diceGrab := DiceGrab{}
		diceSegments := strings.Split(seg, ", ")
		for _, diceSeg := range diceSegments {
			diceGrab = append(diceGrab, parseDices(diceSeg))
		}
		diceGrabs = append(diceGrabs, diceGrab)
	}
	return diceGrabs
}

func parseDices(segment string) Dices {
	segments := strings.Split(segment, " ")
	count, err := strconv.Atoi(segments[0])
	if err != nil {
		log.Error(err)
		panic(err)
	}
	return Dices{
		Count: count,
		Color: segments[1],
	}
}

func getGameIdFromSegment(segment string) int {
	gameId, err := strconv.Atoi(strings.TrimLeft(segment, "Game "))
	if err != nil {
		log.Error(fmt.Sprintf("failed to parse game id: %s", segment))
		panic(errors.New("panic!"))
	}
	return gameId
}

func one(input []Game) {
	const (
		maxRed   = 12
		maxGreen = 13
		maxBlue  = 14
	)

	sum := 0
	for _, game := range input {
		gamePossible := true
		for _, grab := range game.DiceGrabs {
			for _, dice := range grab {
				switch dice.Color {
				case Green:
					if dice.Count > maxGreen {
						gamePossible = false
					}
					break
				case Red:
					if dice.Count > maxRed {
						gamePossible = false
					}
					break
				case Blue:
					if dice.Count > maxBlue {
						gamePossible = false
					}
					break
				default:
					log.Error(fmt.Sprintf("unsupported color found %s", dice.Color))
				}
			}
		}
		if gamePossible {
			sum += game.ID
		}
	}
	log.Info(fmt.Sprintf("part 1 solution: %d", sum))
}

func two(input []Game) {
	sum := 0
	for _, game := range input {
		var (
			minRed   = 0
			minGreen = 0
			minBlue  = 0
		)
		for _, grab := range game.DiceGrabs {
			for _, dice := range grab {
				switch dice.Color {
				case Green:
					if dice.Count > minGreen {
						minGreen = dice.Count
					}
					break
				case Red:
					if dice.Count > minRed {
						minRed = dice.Count
					}
					break
				case Blue:
					if dice.Count > minBlue {
						minBlue = dice.Count
					}
					break
				default:
					log.Error(fmt.Sprintf("unsupported color found %s", dice.Color))
				}
			}
		}
		sum += (minRed * minGreen * minBlue)
	}
	log.Info(fmt.Sprintf("part 2 solution: %d", sum))
}
