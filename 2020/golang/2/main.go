package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type passwordPolicy struct {
	Character string
	Minimum   int
	Maximum   int
	Password  string
}

func (pol *passwordPolicy) IsCorrect() bool {
	count := strings.Count(pol.Password, pol.Character)
	return pol.Minimum <= count && count <= pol.Maximum
}

func (pol *passwordPolicy) IsCorrectTwo() bool {
	containsMinString := string(pol.Password[pol.Minimum]) == pol.Character
	containsMaxString := string(pol.Password[pol.Maximum]) == pol.Character
	return (containsMinString || containsMaxString) && !(containsMinString && containsMaxString)
}

func main() {
	printPasswordInformation(readFile())
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile() []passwordPolicy {
	passwordPolicies := []passwordPolicy{}

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()

		policyAndPassword := strings.Split(line, ":")
		password := policyAndPassword[1]
		policy := strings.Split(policyAndPassword[0], " ")
		policyCharacter := policy[1]
		minmax := strings.Split(policy[0], "-")
		minimum, err := strconv.Atoi(minmax[0])
		check(err)
		maximum, err := strconv.Atoi(minmax[1])
		check(err)

		passwordPolicy := passwordPolicy{
			Character: policyCharacter,
			Minimum:   minimum,
			Maximum:   maximum,
			Password:  password,
		}
		passwordPolicies = append(passwordPolicies, passwordPolicy)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return passwordPolicies
}

func printPasswordInformation(passwords []passwordPolicy) {
	faultyCount := 0
	correctCount := 0
	faultyTwo := 0
	correctTwo := 0
	for _, password := range passwords {
		if password.IsCorrect() {
			correctCount++
		} else {
			faultyCount++
		}

		if password.IsCorrectTwo() {
			correctTwo++
		} else {
			faultyTwo++
		}
	}
	fmt.Printf("Correct passwords: %d \n", correctCount)
	fmt.Printf("Faulty passwords: %d \n\n", faultyCount)

	fmt.Printf("Correct passwords two: %d \n", correctTwo)
	fmt.Printf("Faulty passwords two: %d \n", faultyTwo)
}
