package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	inputNumbers := readFile()

	twoNumbers(inputNumbers)
	threeNumbers(inputNumbers)
}

func readFile() []int {
	inputNumbers := []int{}

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
		inputNumbers = append(inputNumbers, number)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return inputNumbers
}

func twoNumbers(inputNumbers []int) {
	for index, number := range inputNumbers {
		for i := index + 1; i < len(inputNumbers); i++ {
			sum := number + inputNumbers[i]
			if sum == 2020 {
				fmt.Printf("Numbers: %d, %d \n", number, inputNumbers[i])
				fmt.Printf("Sum: %d \n", number+inputNumbers[i])
				fmt.Printf("Multiply: %d \n\n", number*inputNumbers[i])
				break
			}
		}
	}
}

func threeNumbers(inputNumbers []int) {
	for index, number := range inputNumbers {
		for i := index + 1; i < len(inputNumbers); i++ {
			for j := i + 1; j < len(inputNumbers); j++ {
				sum := number + inputNumbers[i] + inputNumbers[j]
				if sum == 2020 {
					fmt.Printf("Numbers: %d, %d, %d \n", number, inputNumbers[i], inputNumbers[j])
					fmt.Printf("Sum: %d \n", number+inputNumbers[i]+inputNumbers[j])
					fmt.Printf("Multiply: %d \n\n", number*inputNumbers[i]*inputNumbers[j])
					break
				}
			}
		}
	}
}
