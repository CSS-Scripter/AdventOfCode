package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Command funny struct
type Command struct {
	action string
	amount int
}

var (
	commands               []Command
	executedCommandIndexes map[int]bool
)

func main() {
	readFile("./input.txt")
	one()
	two()
}

func readFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		commandStrings := strings.Split(line, " ")
		amount, err := strconv.Atoi(commandStrings[1])
		if err != nil {
			panic(err)
		}
		command := Command{
			action: commandStrings[0],
			amount: amount,
		}
		commands = append(commands, command)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func one() {
	executedCommandIndexes := map[int]bool{}
	executingDoubleCode := false
	currentCommandIndex := 0
	accumilator := 0
	for !executingDoubleCode {
		if executedCommandIndexes[currentCommandIndex] == true {
			executingDoubleCode = true
			break
		}
		if currentCommandIndex > len(commands) {
			fmt.Println("Executed!")
			break
		}
		executedCommandIndexes[currentCommandIndex] = true
		command := commands[currentCommandIndex]
		fmt.Printf("%v: %+v\n", currentCommandIndex+1, command)
		switch command.action {
		case "acc":
			accumilator += command.amount
			currentCommandIndex++
			break
		case "jmp":
			currentCommandIndex += command.amount
			break
		case "nop":
			currentCommandIndex++
			break
		}
	}
	fmt.Println(accumilator)
}

func two() {
	changeAction := 1
	executedSuccesfully := false
	for !executedSuccesfully {
		executedSuccesfully = executeTwo(changeAction)
		changeAction++
	}
}

func executeTwo(changeEncounter int) bool {
	executedCommandIndexes := map[int]bool{}
	executingDoubleCode := false
	currentCommandIndex := 0
	accumilator := 0
	encounteredJmpsOrNops := 0
	success := false
	for !executingDoubleCode {
		if currentCommandIndex >= len(commands) {
			fmt.Println("Executed!")
			executingDoubleCode = true
			success = true
			break
		}
		if executedCommandIndexes[currentCommandIndex] == true {
			executingDoubleCode = true
			break
		}
		executedCommandIndexes[currentCommandIndex] = true
		command := commands[currentCommandIndex]
		fmt.Printf("%v: %+v\n", currentCommandIndex+1, command)
		switch command.action {
		case "acc":
			accumilator += command.amount
			currentCommandIndex++
			break
		case "jmp":
			encounteredJmpsOrNops++
			if encounteredJmpsOrNops == changeEncounter {
				currentCommandIndex++
				break
			}
			currentCommandIndex += command.amount
			break
		case "nop":
			encounteredJmpsOrNops++
			if encounteredJmpsOrNops == changeEncounter {
				currentCommandIndex += command.amount
				break
			}
			currentCommandIndex++
			break
		}
	}
	fmt.Println(accumilator)
	return success
}
