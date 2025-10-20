package d04

import (
	"aoc2023/src/util"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
)

type Game struct {
	ID             int
	WinningNumbers []int
	Points         int
}

func NewGame(id int, numbers, luckyNumbers []int) Game {
	winningNumbers := []int{}
	for _, num := range numbers {
		for _, winNum := range luckyNumbers {
			if num == winNum {
				winningNumbers = append(winningNumbers, num)
			}
		}
	}

	var points = 0
	if len(winningNumbers) > 0 {
		points = int(math.Pow(2, float64(len(winningNumbers)-1)))
	}

	return Game{
		ID:             id,
		WinningNumbers: winningNumbers,
		Points:         points,
	}
}

func Main() {
	data, err := util.ReadInput(4)
	if err != nil {
		log.Error(err)
		panic(err)
	}

	games := parseInput(data)
	one(games)
	two(games)
}

func parseInput(data []byte) []Game {
	games := []Game{}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		segments := strings.Split(line, ": ")
		id := getGameId(segments[0])
		winningNumbers, numbers := getNumbersFromLine(segments[1])
		game := NewGame(id, numbers, winningNumbers)
		games = append(games, game)
	}
	return games
}

func getGameId(seg string) int {
	id, err := strconv.Atoi(strings.TrimLeft(seg, "Card "))
	if err != nil {
		log.Error(err)
		panic(err)
	}
	return id
}

func getNumbersFromLine(seg string) ([]int, []int) {
	segments := strings.Split(seg, " | ")
	return getNumbersFromSegment(segments[0]), getNumbersFromSegment(segments[1])
}

func getNumbersFromSegment(seg string) []int {
	numStrings := strings.Split(seg, " ")
	nums := []int{}
	for _, numString := range numStrings {
		if strings.TrimSpace(numString) == "" {
			continue
		}
		num, err := strconv.Atoi(strings.TrimSpace(numString))
		if err != nil {
			log.Error(err)
			panic(err)
		}
		nums = append(nums, num)
	}
	return nums
}

func one(games []Game) {
	var points = 0
	for _, game := range games {
		points += game.Points
	}
	log.Info(fmt.Sprintf("part 1 solution: %d", points))
}

func two(games []Game) {
	gameBacklog := map[int]int{}
	for i := range games {
		gameBacklog[i] = 1
	}
	for i := range games {
		count := gameBacklog[i]
		if count > 0 {
			points := len(games[i].WinningNumbers)
			if points == 0 {
				continue
			}

			var subI int
			for subI = i + 1; subI <= i+points; subI++ {
				gameBacklog[subI] += count
			}
		}
	}

	cards := 0
	for i := range games {
		cards += gameBacklog[i]
	}

	log.Info(fmt.Sprintf("part 2 solution: %d", cards))
}
