package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var (
	mask   string
	memory map[string]int = map[string]int{}
)

func main() {
	one()
	two()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func one() {
	input, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	for _, line := range lines {
		if strings.Contains(line, "mask") {
			re := regexp.MustCompile("(0?1?X?)+")
			maskValue := re.FindAllString(line, -1)
			mask = maskValue[len(maskValue)-1]
		} else {
			re := regexp.MustCompile("[0-9]+")
			memPosVal := re.FindAllString(line, -1)
			memVal, _ := strconv.Atoi(memPosVal[1])
			memory[memPosVal[0]] = passValueThroughMask(memVal)
		}
	}
	total := 0
	for _, mem := range memory {
		total += mem
	}
	fmt.Println(total)
}

func passValueThroughMask(val int) int {
	memory = map[string]int{}
	binVal := fmt.Sprintf("%b", val)
	binVal = addStartingZeros(binVal)
	returnBin := ""
	for i, char := range mask {
		switch char {
		case 'X':
			returnBin += string(binVal[i])
			break
		case '1':
			returnBin += "1"
			break
		case '0':
			returnBin += "0"
			break
		}
	}
	return binToInt(returnBin)
}

func addStartingZeros(bin string) string {
	for len(bin) < len(mask) {
		bin = "0" + bin
	}
	return bin
}

func two() {
	memory = map[string]int{}
	input, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	for _, line := range lines {
		if strings.Contains(line, "mask") {
			re := regexp.MustCompile("(0?1?X?)+")
			maskValue := re.FindAllString(line, -1)
			mask = maskValue[len(maskValue)-1]
		} else {
			re := regexp.MustCompile("[0-9]+")
			memPosVal := re.FindAllString(line, -1)
			memVal, _ := strconv.Atoi(memPosVal[1])
			for _, addr := range getAllAddresses(stringToInt(memPosVal[0])) {
				memory[fmt.Sprint(addr)] = memVal
			}
		}
	}
	total := 0
	for _, mem := range memory {
		total += mem
	}
	fmt.Println(total)
}

func getAllAddresses(addr int) []int {
	binVal := fmt.Sprintf("%b", addr)
	binVal = addStartingZeros(binVal)
	returnBin := ""
	addresses := []int{}
	for i, char := range mask {
		switch char {
		case 'X':
			returnBin += "X"
			break
		case '1':
			returnBin += "1"
			break
		case '0':
			returnBin += string(binVal[i])
			break
		}
	}

	variations := getVariatiesFromMask()
	for v := 0; v < variations; v++ {
		filledMask := placeBinOverMask(returnBin, fmt.Sprintf("%b", v))
		addresses = append(addresses, binToInt(filledMask))
	}
	return addresses
}

func placeBinOverMask(returnMask string, bin string) string {
	bIndex := 0
	for len(bin) < countXInMask(returnMask) {
		bin = "0" + bin
	}
	for strings.Contains(returnMask, "X") {
		returnMask = strings.Replace(returnMask, "X", string(bin[bIndex]), 1)
		bIndex++
	}
	return returnMask
}

func getVariatiesFromMask() int {
	returnVal := int(math.Pow(2, float64(countXInMask(mask))))
	return returnVal
}

func countXInMask(pmask string) int {
	xCount := 0
	for _, run := range pmask {
		if run == 'X' {
			xCount++
		}
	}
	return xCount
}

func binToInt(bin string) int {
	num, err := strconv.ParseInt(bin, 2, 64)
	check(err)
	return int(num)
}

func stringToInt(str string) int {
	num, err := strconv.Atoi(str)
	check(err)
	return num
}
