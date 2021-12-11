package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var (
	binary []string
)

func main() {
	one()
	two()
}

func one() {
	prepareInput("input.txt")
	gamma := ""
	epsilon := ""
	for i := 0; i < len(binary[0]); i++ {
		counts := map[byte]int{
			'0': 0,
			'1': 0,
		}
		for _, bin := range binary {
			counts[bin[i]] += 1
		}
		if counts['1'] > counts['0'] {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}
	gammaInt, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonInt, _ := strconv.ParseInt(epsilon, 2, 64)
	fmt.Printf("One: %d\n", gammaInt*epsilonInt)
}

func two() {
	prepareInput("input.txt")
	oxygenBin := binary
	scrubberBin := binary
	for i := 0; i < len(binary[0]); i++ {
		counts := map[string]map[byte]int{
			"oxygen":   {'0': 0, '1': 0},
			"scrubber": {'0': 0, '1': 0},
		}
		if len(oxygenBin) > 1 {
			for _, o2bin := range oxygenBin {
				counts["oxygen"][o2bin[i]] += 1
			}
			if counts["oxygen"]['1'] >= counts["oxygen"]['0'] {
				oxygenBin = filter(oxygenBin, i, '1')
			} else {
				oxygenBin = filter(oxygenBin, i, '0')
			}
		}

		if len(scrubberBin) > 1 {
			for _, scrubbin := range scrubberBin {
				counts["scrubber"][scrubbin[i]] += 1
			}
			if counts["scrubber"]['0'] <= counts["scrubber"]['1'] {
				scrubberBin = filter(scrubberBin, i, '0')
			} else {
				scrubberBin = filter(scrubberBin, i, '1')
			}
		}
	}
	oxygen, _ := strconv.ParseInt(oxygenBin[0], 2, 64)
	scrubber, _ := strconv.ParseInt(scrubberBin[0], 2, 64)
	fmt.Printf("Two: %d\n", oxygen*scrubber)
}

func filter(binArr []string, i int, binChar byte) []string {
	filtered := []string{}
	for _, bin := range binArr {
		if string(bin[i]) == string(binChar) {
			filtered = append(filtered, bin)
		}
	}
	return filtered
}

func prepareInput(file string) {
	binary = []string{}
	input, _ := ioutil.ReadFile(file)
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	for _, line := range lines {
		binary = append(binary, line)
	}
}
