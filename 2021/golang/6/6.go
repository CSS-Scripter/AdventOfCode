package main

import (
	"io/ioutil"
	"strings"
	"fmt"
	"strconv"
)

var (
	lanternFish [9]int = [9]int{}
)


func main() {
	prepareInput("input.txt")
	simulateDays(80)
	countLanternFish()
}

func one() {
	prepareInput("input.txt")
	simulateDays(80)
	countLanternFish()
}

func two() {
	prepareInput("input.txt")
	simulateDays(256)
	countLanternFish()
}

func prepareInput(file string) {
	input, _ := ioutil.ReadFile(file)
	numstrings := strings.Split(strings.TrimSpace(string(input)), ",")
	for _, numstring := range numstrings {
		num, _ := strconv.Atoi(numstring)
		lanternFish[num] += 1
	}
}

func simulateDays(days int) {
	for day := 0; day < days; day++ {
		amountToReproduce := lanternFish[0]
		for i, lantern := range lanternFish[1:] {
			lanternFish[i] = lantern
		}
		lanternFish[6] += amountToReproduce
		lanternFish[8] = amountToReproduce
	}
}

func countLanternFish() {
	sum := 0
	for _, amount := range lanternFish {
		sum += amount
	}
	fmt.Println(sum)
}