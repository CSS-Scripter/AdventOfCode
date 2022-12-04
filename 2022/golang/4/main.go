package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Range struct {
	Min int
	Max int
}

type RangePair struct {
	Range1 Range
	Range2 Range
}

var RangePairs []RangePair

func (r *Range) IsContainedBy(rr Range) bool {
	return r.Min <= rr.Min && r.Max >= rr.Max
}

func (r *Range) IsOverlappedBy(rr Range) bool {
	return (r.Max >= rr.Min && r.Min <= rr.Max) || (rr.Max >= r.Min && rr.Min <= r.Max)
}

func (r *RangePair) HasContainingPair() bool {
	return r.Range1.IsContainedBy(r.Range2) || r.Range2.IsContainedBy(r.Range1)
}

func (r *RangePair) HasOverlappingPair() bool {
	return r.Range1.IsOverlappedBy(r.Range2)
}

func main() {
	prepareInput("input.txt")
	one()
	two()
}

func one() {
	count := 0
	for _, pair := range RangePairs {
		if pair.HasContainingPair() {
			count++
		}
	}
	fmt.Printf("One: %d\n", count)
}

func two() {
	count := 0
	for _, pair := range RangePairs {
		if pair.HasOverlappingPair() {
			fmt.Println(pair)
			count++
		}
	}
	fmt.Printf("One: %d\n", count)
}

func prepareInput(file string) {
	RangePairs = []RangePair{}
	input, _ := ioutil.ReadFile(file)
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	for _, line := range lines {
		ranges := strings.Split(line, ",")
		RangePairs = append(RangePairs, RangePair{
			parseRangeFromString(ranges[0]),
			parseRangeFromString(ranges[1]),
		})
	}
}

func parseRangeFromString(line string) Range {
	minMax := strings.Split(line, "-")
	Min, _ := strconv.Atoi(minMax[0])
	Max, _ := strconv.Atoi(minMax[1])
	return Range{Min, Max}
}
