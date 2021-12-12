package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

var (
	Caves map[string]*Cave
)

func InitCaveSystem() {
	Caves = map[string]*Cave{
		"start": {
			Name:        "start",
			IsBig:       false,
			ConnectedTo: []*Cave{},
		},
		"end": {
			Name:        "end",
			IsBig:       false,
			ConnectedTo: []*Cave{},
		},
	}
}

func Link(from, to string) {
	_, fromExists := Caves[from]
	_, toExists := Caves[to]
	if !fromExists {
		CreateCave(from)
	}
	if !toExists {
		CreateCave(to)
	}
	Caves[from].ConnectedTo = append(Caves[from].ConnectedTo, Caves[to])
	Caves[to].ConnectedTo = append(Caves[to].ConnectedTo, Caves[from])
}

func CreateCave(name string) {
	Caves[name] = &Cave{
		Name:        name,
		IsBig:       IsUpperCase(name),
		ConnectedTo: []*Cave{},
	}
}

func IsUpperCase(s string) bool {
	for _, l := range s {
		if !unicode.IsUpper(l) && unicode.IsLetter(l) {
			return false
		}
	}
	return true
}

type Cave struct {
	Name        string
	ConnectedTo []*Cave
	IsBig       bool
}

func main() {
	one()
	two()
}

func one() {
	prepareInput("example_input.txt")
	paths, success := Visit(Caves["start"], map[string]bool{}, true)
	if !success {
		panic("Mission failed! We'll get 'em next time")
	}
	fmt.Printf("One: %d\n", len(paths))
}

func two() {
	prepareInput("input.txt")
	paths, success := Visit(Caves["start"], map[string]bool{}, false)
	if !success {
		panic("Mission failed! We'll get 'em next time")
	}
	fmt.Printf("Two: %d\n", len(paths))
}

func prepareInput(file string) {
	InitCaveSystem()
	input, _ := ioutil.ReadFile(file)
	inputStrings := strings.Split(strings.TrimSpace(string(input)), "\n")
	for _, inputString := range inputStrings {
		from := strings.Split(inputString, "-")[0]
		to := strings.Split(inputString, "-")[1]
		Link(from, to)
	}
}

func Visit(c *Cave, visited map[string]bool, hasSmallCaveBeenVisited bool) ([]string, bool) {
	visited[c.Name] = true
	if c.Name == "end" {
		return []string{"end"}, true
	}

	visits := []string{}

	for _, newCave := range c.ConnectedTo {
		if newCave.Name == "start" {
			continue
		}
		_, hasBeenVisited := visited[newCave.Name]
		if newCave.IsBig || !hasBeenVisited || (hasBeenVisited && !hasSmallCaveBeenVisited && !newCave.IsBig) {
			newVisitSmallCave := hasSmallCaveBeenVisited || (hasBeenVisited && !newCave.IsBig)
			paths, valid := Visit(newCave, CopyVisitedMap(visited), newVisitSmallCave)
			if valid {
				for _, path := range paths {
					visits = append(visits, path)
				}
			}
		} else {
			continue
		}
	}
	if len(visits) == 0 {
		return visits, false
	}
	for i, path := range visits {
		visits[i] = c.Name + "," + path
	}
	return visits, true
}

func CopyVisitedMap(visitedMap map[string]bool) map[string]bool {
	newMap := map[string]bool{}
	for name, visited := range visitedMap {
		newMap[name] = visited
	}
	return newMap
}
