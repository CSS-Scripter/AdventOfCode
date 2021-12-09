package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	one()
	two()
}

func one() {
	questionSums := []int{}
	input := readFile()
	groupsQuestions := map[rune]bool{}
	for _, line := range input {
		if line != "" {
			for _, questionChar := range line {
				groupsQuestions[questionChar] = true
			}
		} else {
			questionSums = append(questionSums, len(groupsQuestions))
			groupsQuestions = map[rune]bool{}
		}
	}
	questionSums = append(questionSums, len(groupsQuestions))
	groupsQuestions = map[rune]bool{}

	fmt.Printf("Part one: %v\n", sumCounts(questionSums))
}

func two() {
	total := 0
	input := readFile()
	group := []map[rune]bool{}
	for _, line := range input {
		person := map[rune]bool{}
		if line != "" {
			for _, questionChar := range line {
				person[questionChar] = true
			}
			group = append(group, person)
		} else {
			total += getTotalFromGroup(group)
			group = []map[rune]bool{}
		}
	}
	total += getTotalFromGroup(group)

	fmt.Printf("Part two: %v\n", total)
}

func getTotalFromGroup(group []map[rune]bool) int {
	total := 0
	questionsAnsweredAmount := map[rune]int{}
	for _, person := range group {
		for question := range person {
			questionsAnsweredAmount[question]++
		}
	}
	for _, amount := range questionsAnsweredAmount {
		if amount == len(group) {
			total++
		}
	}
	return total
}

func readFile() []string {
	returnValue := []string{}

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		returnValue = append(returnValue, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return returnValue
}

func sumCounts(counts []int) int {
	total := 0
	for _, count := range counts {
		total += count
	}
	return total
}
