package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var (
	intCodeProgram []int
	executables    map[int]Executable = map[int]Executable{
		1:  Addition,
		2:  Multiplication,
		99: Terminate,
	}
	running = true
)

type Executable func(int, int, int)

func main() {
	one()
	two()
}

func one() {
	prepareInput("input.txt")
	intCodeProgram[1] = 12
	intCodeProgram[2] = 2
	for i := 0; i < len(intCodeProgram); i += 4 {
		executables[intCodeProgram[i]](i+1, i+2, i+3)
		if !running {
			break
		}
	}
	fmt.Printf("One: %d\n", intCodeProgram[0])
}

func two() {
	prepareInput("input.txt")
	noun, verb := getNounAndVerb()
	fmt.Printf("Two: %d\n", (100*noun)+verb)
}

func getNounAndVerb() (int, int) {
	intCodeBackup := make([]int, len(intCodeProgram))
	copy(intCodeBackup, intCodeProgram)
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			copy(intCodeProgram, intCodeBackup)
			intCodeProgram[1] = noun
			intCodeProgram[2] = verb
			running = true
			for i := 0; i < len(intCodeProgram); i += 4 {
				executables[intCodeProgram[i]](i+1, i+2, i+3)
				if !running {
					break
				}
			}
			if intCodeProgram[0] == 19690720 {
				return noun, verb
			}
		}
	}
	return 0, 0
}

func prepareInput(file string) {
	intCodeProgram = []int{}
	input, _ := ioutil.ReadFile(file)
	inputStrings := strings.Split(strings.TrimSpace(string(input)), ",")
	for _, inputLine := range inputStrings {
		intcode, _ := strconv.Atoi(inputLine)
		intCodeProgram = append(intCodeProgram, intcode)
	}
}

func Addition(pos1, pos2, out int) {
	if !IsPosValid(pos1) || !IsPosValid(pos2) || !IsPosValid(out) {
		running = false
		return
	}
	intCodeProgram[intCodeProgram[out]] = intCodeProgram[intCodeProgram[pos1]] + intCodeProgram[intCodeProgram[pos2]]
}

func Multiplication(pos1, pos2, out int) {
	if !IsPosValid(pos1) || !IsPosValid(pos2) || !IsPosValid(out) {
		running = false
		return
	}
	intCodeProgram[intCodeProgram[out]] = intCodeProgram[intCodeProgram[pos1]] * intCodeProgram[intCodeProgram[pos2]]
}

func IsPosValid(pos int) bool {
	return len(intCodeProgram) > pos && len(intCodeProgram) > intCodeProgram[pos]
}

func Terminate(_, _, _ int) {
	running = false
}
