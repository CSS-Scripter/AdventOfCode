package main

import (
	"io/ioutil"
	"strings"
	"fmt"
	"sort"
)

var (
	lines []Line
	normalBracket Bracket = Bracket{
		OpeningBracket: '(',
		ClosingBracket: ')',
		Score: 3,
		ClosingScore: 1,
	}
	blockBracket Bracket = Bracket{
		OpeningBracket: '[',
		ClosingBracket: ']',
		Score: 57,
		ClosingScore: 2,
	}
	curlyBracket Bracket = Bracket{
		OpeningBracket: '{',
		ClosingBracket: '}',
		Score: 1197,
		ClosingScore: 3,
	}
	arrowBracket Bracket = Bracket{
		OpeningBracket: '<',
		ClosingBracket: '>',
		Score: 25137,
		ClosingScore: 4,
	}
	brackets map[byte]Bracket = map[byte]Bracket{
		'(': normalBracket,
		')': normalBracket,
		'[': blockBracket,
		']': blockBracket,
		'{': curlyBracket,
		'}': curlyBracket,
		'<': arrowBracket,
		'>': arrowBracket,
	}
)

type Stack []byte

func (s Stack) Peek() byte {
	if len(s) > 0 { return s[len(s)-1] }
	return 0
}

func (s *Stack) Pop() byte {
	if len(*s) > 0 {
		lastElement := s.Peek()
		*s = (*s)[:len(*s)-1]
		return lastElement
	}
	return 0
}

func (s *Stack) Push(e byte) {
	*s = append(*s, e)
}

func (s Stack) IsEmpty() bool {
	return len(s) == 0
}

func (s Stack) String() string {
	return string(s)
}

type Bracket struct {
	OpeningBracket byte
	ClosingBracket byte
	Score int
	ClosingScore int
}

type Line struct {
	Brackets []byte
	BracketStack *Stack
}

func (l Line) isCorrupt() int {
	for _, b := range l.Brackets {
		bracketType := brackets[b]
		if b == bracketType.OpeningBracket {
			l.BracketStack.Push(b)
			continue
		}
		lastOpenedBracketType := brackets[l.BracketStack.Peek()]
		if (lastOpenedBracketType.ClosingBracket != b) {
			fmt.Printf("Expected %s, but found %s instead.\n", string(lastOpenedBracketType.ClosingBracket), string(b))
			return bracketType.Score
		}
		l.BracketStack.Pop()
	}
	return 0
}

func main() {
	prepareInput("input.txt")
	scores := []int{}
	corruptScore := 0
	for _, l := range lines {
		corruptLScore := l.isCorrupt()
		if corruptLScore > 0 { corruptScore += corruptLScore; continue }
		lineScore := 0
		fmt.Print("Complete by adding ")
		for !l.BracketStack.IsEmpty()  {
			bType := brackets[l.BracketStack.Pop()]
			lineScore *= 5
			lineScore += bType.ClosingScore
			fmt.Print(string(bType.ClosingBracket))
		}
		fmt.Printf(" - %d total points\n", lineScore)
		scores = append(scores, lineScore)
	}
	sort.Ints(scores)
	fmt.Printf("One: %d\nTwo: %d\n", corruptScore, scores[len(scores) / 2])
}

func prepareInput(file string) {
	input, _ := ioutil.ReadFile(file)
	inputStrings := strings.Split(strings.TrimSpace(string(input)), "\n")
	for _, inputLine := range inputStrings {
		brackets := []byte(inputLine)
		line := Line{Brackets: brackets, BracketStack: &Stack{}}
		lines = append(lines, line)
	}
}