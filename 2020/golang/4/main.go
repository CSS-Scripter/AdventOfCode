package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Passport haha funny struct
type Passport struct {
	byr int
	iyr int
	eyr int
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

// IsValid checks the validity of the passport
func (pass *Passport) IsValid() bool {
	isValid := true
	isValid = pass.isByrValid() && isValid
	isValid = pass.isIyrValid() && isValid
	isValid = pass.isEyrValid() && isValid
	isValid = pass.isHgtValid() && isValid
	isValid = pass.isHclValid() && isValid
	isValid = pass.isEclValid() && isValid
	isValid = pass.isPidValid() && isValid
	return isValid
}

func (pass *Passport) isByrValid() bool {
	return 1920 <= pass.byr && pass.byr <= 2002
}

func (pass *Passport) isIyrValid() bool {
	return 2010 <= pass.iyr && pass.iyr <= 2020
}

func (pass *Passport) isEyrValid() bool {
	return 2020 <= pass.eyr && pass.eyr <= 2030
}

func (pass *Passport) isHgtValid() bool {
	if matchesRegex("[0-9]+(cm|in)", pass.hgt) {
		regex := regexp.MustCompile("[0-9]+")
		matches := regex.FindAllString(pass.hgt, -1)
		length := stringToInt(matches[0])
		if strings.Contains(pass.hgt, "cm") {
			return 150 <= length && length <= 193
		}
		if strings.Contains(pass.hgt, "in") {
			return 59 <= length && length <= 76
		}
	}
	return false
}

func (pass *Passport) isHclValid() bool {
	return matchesRegex("#([a-fA-F0-9]{6})", pass.hcl)
}

func (pass *Passport) isEclValid() bool {
	validColors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, color := range validColors {
		if color == pass.ecl {
			return true
		}
	}
	return false
}

func (pass *Passport) isPidValid() bool {
	return matchesRegex("[0-9]{9}", pass.pid)
}

func (pass *Passport) String() string {
	return fmt.Sprintf("Passport<byr: %v, iyr: %v, eyr: %v, hgt: %v, hcl: %v, ecl: %v, pid: %v, cid: %v>", pass.byr, pass.iyr, pass.eyr, pass.hgt, pass.hcl, pass.ecl, pass.pid, pass.cid)
}

func matchesRegex(regex string, value string) bool {
	matches, err := regexp.MatchString(regex, value)
	if err != nil {
		panic(err)
	}
	return matches
}

// UnknownFieldError err
type UnknownFieldError struct{}

func (err *UnknownFieldError) Error() string {
	return "Unknown field found!"
}

func main() {
	passports := readFile()
	validCount := 0
	for _, passport := range passports {
		if passport.IsValid() {
			validCount++
		}
	}
	fmt.Printf("%v out of %v passports are valid!\n", validCount, len(passports))
}

func readFile() []Passport {
	returnValue := []Passport{}

	fileIO, err := os.OpenFile("input.txt", os.O_RDWR, 0600)
	if err != nil {
		panic(err)
	}
	defer fileIO.Close()
	rawBytes, err := ioutil.ReadAll(fileIO)
	if err != nil {
		panic(err)
	}

	rawPassportStrings := []string{}
	lines := strings.Split(string(rawBytes), "\n")
	rawPassportString := ""
	for _, line := range lines {
		if line != "" {
			rawPassportString += fmt.Sprintf("%v ", line)
		} else {
			rawPassportStrings = append(rawPassportStrings, rawPassportString)
			rawPassportString = ""
		}
	}
	for _, passString := range rawPassportStrings {
		returnValue = append(returnValue, parseStringToPassport(passString))
	}
	return returnValue
}

func parseStringToPassport(passportString string) Passport {
	passport := Passport{}
	passportString = strings.TrimSpace(passportString)
	for _, passportElement := range strings.Split(passportString, " ") {
		splitElements := strings.Split(passportElement, ":")
		key := splitElements[0]
		value := splitElements[1]
		switch key {
		case "byr":
			passport.byr = stringToInt(value)
			break
		case "iyr":
			passport.iyr = stringToInt(value)
			break
		case "eyr":
			passport.eyr = stringToInt(value)
			break
		case "hgt":
			passport.hgt = value
			break
		case "hcl":
			passport.hcl = value
			break
		case "ecl":
			passport.ecl = value
			break
		case "pid":
			passport.pid = value
			break
		case "cid":
			passport.cid = value
			break
		default:
			panic(new(UnknownFieldError))
		}
	}
	return passport
}

func stringToInt(in string) int {
	out, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}
	return out
}
