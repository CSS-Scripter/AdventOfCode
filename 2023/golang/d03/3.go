package d03

import (
	"aoc2023/src/util"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
)

type Part struct {
	Value              int
	AdjecantCharacters []byte
	IsMachinePart      bool
	PositionY          int
	PositionMinX       int
	PositionMaxX       int
}

func Main() {
	input := parseInput()
	parts := findParts(input)
	one(parts)
	two(input, parts)
}

func parseInput() []string {
	input, err := util.ReadInput(3)
	if err != nil {
		log.Error(err)
		panic(err)
	}

	lines := strings.Split(string(input), "\n")
	return lines
}

func one(parts []Part) {
	sum := 0
	for _, part := range parts {
		if part.IsMachinePart {
			sum += part.Value
		}
	}
	log.Info(fmt.Sprintf("part 1 solution: %d", sum))
}

type Position [2]int

func two(input []string, parts []Part) {
	gearPositions := []Position{}
	for y, line := range input {
		for x, char := range line {
			if char == '*' {
				gearPositions = append(gearPositions, Position{x, y})
			}
		}
	}

	gearParts := []Part{}
	for _, part := range parts {
		if strings.Contains(string(part.AdjecantCharacters), "*") {
			gearParts = append(gearParts, part)
		}
	}

	sum := 0
	for _, pos := range gearPositions {
		x := pos[0]
		y := pos[1]
		values := findGearValues(gearParts, x, y)
		if len(values) == 2 {
			sum += values[0] * values[1]
		}
	}
	log.Info(fmt.Sprintf("part 2 solution: %d", sum))
}

func findParts(input []string) []Part {
	parts := []Part{}

	for y, line := range input {
		var isInNumberSequence = false
		for x, char := range []byte(line) {
			if !isNumber(char) {
				isInNumberSequence = false
				continue
			}

			if isNumber(char) && !isInNumberSequence {
				isInNumberSequence = true
				valueChars := []byte{byte(char)}
				minX := int(math.Max(float64(x-1), 0))
				minY := int(math.Max(float64(y-1), 0))

				posCheck := x + 1
				for posCheck < len(line) && isNumber(line[posCheck]) {
					valueChars = append(valueChars, line[posCheck])
					posCheck++
				}
				maxX := int(math.Min(float64(len(line)-1), float64(posCheck)))
				maxY := int(math.Min(float64(len(input)-1), float64(y+1)))

				surroundingChars := ""
				surroundingChars += input[minY][minX : maxX+1]
				surroundingChars += input[maxY][minX : maxX+1]
				surroundingChars += string(input[y][minX])
				surroundingChars += string(input[y][maxX])
				surroundingChars = filterSurroundChars(surroundingChars)

				value, err := strconv.Atoi(string(valueChars))
				if err != nil {
					log.Error(fmt.Sprintf("failed to parse %s", string(valueChars)))
					panic(err)
				}
				part := Part{
					Value:              value,
					AdjecantCharacters: []byte(surroundingChars),
					IsMachinePart:      len(surroundingChars) > 0,
					PositionMinX:       x,
					PositionMaxX:       maxX - 1,
					PositionY:          y,
				}
				parts = append(parts, part)
				continue
			}
		}
	}
	return parts
}

func isNumber(char byte) bool {
	return 48 <= char && char <= 57
}

func filterSurroundChars(surroundChars string) string {
	re := regexp.MustCompile(`[\.\d]`)
	return re.ReplaceAllString(surroundChars, "")
}

func findGearValues(parts []Part, x, y int) []int {
	minX := x - 1
	maxX := x + 1
	minY := y - 1
	maxY := y + 1

	foundValues := map[int]struct{}{}
	for _, part := range parts {
		if part.PositionY >= minY && part.PositionY <= maxY {
			if (part.PositionMinX >= minX && part.PositionMinX <= maxX) ||
				(part.PositionMaxX >= minX && part.PositionMaxX <= maxX) {
				foundValues[part.Value] = struct{}{}
			}
		}
	}

	values := []int{}
	for val := range foundValues {
		values = append(values, val)
	}

	return values
}
