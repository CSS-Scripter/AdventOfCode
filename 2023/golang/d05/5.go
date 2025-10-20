package d05

import (
	"aoc2023/src/util"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
)

type Map struct {
	Name     string
	Mappings []Mapping
}

type Mapping struct {
	Destination uint64
	Source      uint64
	Size        uint64
}

func (m Mapping) IsNumberInRange(num uint64) bool {
	return num >= m.Source && num < (m.Source+m.Size)
}

func (m Mapping) ReverseIsNumberInRange(num uint64) bool {
	return num >= m.Destination && num < (m.Destination+m.Size)
}

func (m Mapping) RemapNumber(num uint64) uint64 {
	diff := num - m.Source
	return m.Destination + diff
}

func (m Mapping) ReverseRemapNumber(num uint64) uint64 {
	diff := num - m.Destination
	return m.Source + diff
}

func Main() {
	data, err := util.ReadInput(5)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	seeds, maps := ParseInput(data)
	one(seeds, maps)
	two(seeds, maps)
}

func one(seeds []uint64, maps []Map) {
	for _, seedMap := range maps {
		newSeeds := []uint64{}
		for _, seed := range seeds {
			isRemapped := false
			for _, mapping := range seedMap.Mappings {
				if mapping.IsNumberInRange(seed) {
					newSeeds = append(newSeeds, mapping.RemapNumber(seed))
					isRemapped = true
				}
			}
			if !isRemapped {
				newSeeds = append(newSeeds, seed)
			}
		}
		seeds = newSeeds
	}

	smallest := seeds[0]
	for _, seed := range seeds {
		if seed < smallest {
			smallest = seed
		}
	}
	log.Info(fmt.Sprintf("part 1 solution: %d", smallest))
}

func two(seeds []uint64, maps []Map) {
	sort.Slice(maps[0].Mappings[:], func(i, j int) bool {
		return maps[0].Mappings[i].Destination < maps[0].Mappings[j].Destination
	})

	var step uint64 = 1000

	var reverseSeed uint64
	for reverseSeed = 0; true; reverseSeed += step {
		subseed := reverseSeed
		for i := len(maps) - 1; i >= 0; i-- {
			currMap := maps[i]
			for _, mapping := range currMap.Mappings {
				if mapping.ReverseIsNumberInRange(subseed) {
					subseed = mapping.ReverseRemapNumber(subseed)
					break
				}
			}
		}
		isNumValid := isNumInSeedRange(seeds, subseed)
		if isNumValid {
			break
		}
	}

	firstSolve := reverseSeed

	for reverseSeed = firstSolve - step; true; reverseSeed++ {
		subseed := reverseSeed
		for i := len(maps) - 1; i >= 0; i-- {
			currMap := maps[i]
			for _, mapping := range currMap.Mappings {
				if mapping.ReverseIsNumberInRange(subseed) {
					subseed = mapping.ReverseRemapNumber(subseed)
					break
				}
			}
		}
		isNumValid := isNumInSeedRange(seeds, subseed)
		if isNumValid {
			break
		}
	}

	log.Info(fmt.Sprintf("part 2 solution: %d", reverseSeed))
}

func isNumInSeedRange(seeds []uint64, num uint64) bool {
	for i := 0; i < len(seeds); i += 2 {
		if num >= seeds[i] && num < seeds[i]+seeds[i+1] {
			return true
		}
	}
	return false
}

func ParseInput(data []byte) ([]uint64, []Map) {
	sections := strings.Split(string(data), "\n\n")
	seedSection := sections[0]
	mapSections := sections[1:]

	seeds := ParseSeeds(seedSection)
	maps := []Map{}

	for _, mapSection := range mapSections {
		maps = append(maps, ParseMap(mapSection))
	}

	return seeds, maps
}

func ParseSeeds(data string) []uint64 {
	seedStrings := strings.Split(strings.TrimLeft(data, "seeds: "), " ")
	seeds := []uint64{}
	for _, seedString := range seedStrings {
		seed, err := strconv.Atoi(seedString)
		if err != nil {
			log.Error("failed to parse seed", "seedstr", seedString, "err", err)
			panic(err)
		}
		seeds = append(seeds, uint64(seed))
	}
	return seeds
}

func ParseMap(data string) Map {
	lines := strings.Split(data, "\n")
	name := lines[0]
	mappingLines := lines[1:]

	mappings := []Mapping{}
	for _, mappingLine := range mappingLines {
		mappings = append(mappings, ParseMapping(mappingLine))
	}
	return Map{
		Name:     name,
		Mappings: mappings,
	}
}

func ParseMapping(line string) Mapping {
	values := strings.Split(strings.TrimSpace(line), " ")
	parsedValues := make([]uint64, 3)
	for i, val := range values {
		parsedVal, err := strconv.Atoi(val)
		if err != nil {
			log.Error("failed to parse number", "numstring", val)
			panic(err)
		}
		parsedValues[i] = uint64(parsedVal)
	}
	return Mapping{parsedValues[0], parsedValues[1], parsedValues[2]}
}
