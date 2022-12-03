package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Rucksack struct {
	Comp1 string
	Comp2 string
}

var characterScores = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func (r *Rucksack) String() string {
	return fmt.Sprintf("Comp1: %v\nComp2: %v\n\n", string(r.Comp1), string(r.Comp2))
}

func (r *Rucksack) GetAllContents() string {
	return r.Comp1 + r.Comp2
}

func (r *Rucksack) FindCommonItem() string {
	for _, byte1 := range r.Comp1 {
		for _, byte2 := range r.Comp2 {
			if byte1 == byte2 {
				return string(byte1)
			}
		}
	}
	panic("write better alghoritm pleb")
}

var rucksacks = []Rucksack{}

func prepareInput(file string) {
	input, _ := ioutil.ReadFile(file)
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	for _, line := range lines {
		rucksacks = append(rucksacks,
			Rucksack{
				Comp1: line[:len(line)/2],
				Comp2: line[len(line)/2:],
			},
		)
	}
}

func main() {
	prepareInput("input.txt")
	one()
	two()
}

func one() {
	score := 0
	for _, r := range rucksacks {
		commonItem := r.FindCommonItem()
		score += getScoreForRune(byte(commonItem[0]))
	}
	fmt.Println(score)
}

func two() {
	score := 0
	for i := 0; i < len(rucksacks); i += 3 {
		score += findCommonItemAmongThreeRucksacks(rucksacks[i], rucksacks[i+1], rucksacks[i+2])
	}
	fmt.Println(score)
}

func getScoreForRune(in byte) int {
	for i, comp := range characterScores {
		if in == comp {
			return i + 1
		}
	}
	return -1
}

func findCommonItemAmongThreeRucksacks(r1, r2, r3 Rucksack) int {
	r1c := []byte(r1.GetAllContents())
	r2c := []byte(r2.GetAllContents())
	r3c := []byte(r3.GetAllContents())
	for _, rr1 := range r1c {
		for _, rr2 := range r2c {
			for _, rr3 := range r3c {
				if rr1 == rr2 && rr2 == rr3 {
					return getScoreForRune(rr1)
				}
			}
		}
	}
	panic("write better alghoritm pleb")
}
