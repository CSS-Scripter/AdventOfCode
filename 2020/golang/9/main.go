package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var input []int

func main() {
	readFile()
	target := 0
	for i := 25; i < len(input); i++ {
		target = one(input[i-25:i], input[i])
		if target != 0 {
			break
		}
	}
	two(target)
}

func readFile() {
	input = []int{}
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		input = append(input, number)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func one(inputNumbers []int, target int) int {
	for index, number := range inputNumbers {
		for i := index + 1; i < len(inputNumbers); i++ {
			sum := number + inputNumbers[i]
			if sum == target {
				return 0
			}
		}
	}
	fmt.Printf("There is no match found for number: %v\n", target)
	return target
}

func two(target int) {
	for startingIndex := range input {
		smallest := input[startingIndex]
		largest := smallest
		sum := smallest
		numbers := []int{smallest}
		for i := startingIndex + 1; i < len(input); i++ {
			sum += input[i]
			numbers = append(numbers, input[i])
			if input[i] > largest {
				largest = input[i]
			}
			if input[i] < smallest {
				smallest = input[i]
			}
			if sum == target {
				fmt.Printf("Found array of numbers is %v, summing to %v\n", numbers, target)
				fmt.Printf("Sum of smallest and largest numbers is %v\n", smallest+largest)
				break
			}
			if sum > target {
				break
			}
		}
	}
}
