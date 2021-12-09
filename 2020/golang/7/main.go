package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var bags map[string]map[string]int

func main() {
	parseLinesIntoBags(readFile())
	partOne()
	partTwo()
}

func parseLinesIntoBags(lines []string) {
	bags = map[string]map[string]int{}
	for _, line := range lines {
		mainBag := line[:(strings.Index(line, "contain") - 1)]
		mainBag = cleanupBagName(mainBag)
		r, _ := regexp.Compile("[0-9]+ [a-z ]+ bag")
		subBagStrings := r.FindAllString(line, -1)
		bags[mainBag] = map[string]int{}
		for _, bagString := range subBagStrings {
			subBagName := bagString[(strings.Index(bagString, " ") + 1):]
			subBagName = cleanupBagName(subBagName)
			subBagAmountString := bagString[:(strings.Index(bagString, " "))]
			subBagAmount, _ := strconv.Atoi(subBagAmountString)
			bags[mainBag][subBagName] = subBagAmount
		}
	}
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

func cleanupBagName(bagName string) string {
	bagName = strings.TrimSuffix(bagName, " bags")
	bagName = strings.TrimSuffix(bagName, " bag")
	return bagName
}

func partOne() {
	totalBagsContainingGoldBags := 0
	for key := range bags {
		if canBagContainGoldenBag(key) {
			totalBagsContainingGoldBags++
		}
	}
	fmt.Printf("There are %v bags that can carry atleast 1 shiny gold bag.\n", totalBagsContainingGoldBags)
}

func canBagContainGoldenBag(bag string) bool {
	canContain := false
	for key := range bags[bag] {
		if key == "shiny gold" {
			return true
		}
		canContain = canContain || canBagContainGoldenBag(key)
	}
	return canContain
}

func partTwo() {
	amountOfBagsInsideGoldBag := countBagsInsideBag("shiny gold")
	fmt.Printf("There are %v bags inside the shiny gold bag.\n", amountOfBagsInsideGoldBag)
}

func countBagsInsideBag(bag string) int {
	total := 0
	for key, amount := range bags[bag] {
		total += amount * (1 + countBagsInsideBag(key))
	}
	return total
}
