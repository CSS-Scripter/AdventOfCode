package d08

import (
	"aoc2023/src/util"
	"fmt"
	"regexp"
	"strings"

	"github.com/charmbracelet/log"
)

type Instructions string

type Route struct {
	Name  string
	Left  *Route
	Right *Route
}

func Main() {
	data, err := util.ReadInput(8)
	if err != nil {
		log.Error("failed to read input")
		panic(err)
	}

	one(ParseInput(data))
	two(ParseInputTwo(data))
}

func one(instructions string, root *Route) {
	currentRoute := root
	solved := false
	step := 0

	for solved == false {
		for _, action := range instructions {
			step++
			if action == 'L' {
				currentRoute = currentRoute.Left
			} else {
				currentRoute = currentRoute.Right
			}
			if currentRoute.Name == "ZZZ" {
				solved = true
				log.Info(fmt.Sprintf("part 1 solution %d", step))
				break
			}
		}
	}
}

func two(instructions string, routes []*Route) {
	stepsRequired := []int{}
	for _, route := range routes {
		solved := false
		step := 0
		currentRoute := route
		for solved == false {
			for _, action := range instructions {
				step++
				if action == 'L' {
					currentRoute = currentRoute.Left
				} else {
					currentRoute = currentRoute.Right
				}
				if currentRoute.Name[len(currentRoute.Name)-1] == 'Z' {
					solved = true
					stepsRequired = append(stepsRequired, step)
					break
				}
			}
		}
	}

	result := LCM(stepsRequired[0], stepsRequired[1], stepsRequired[2:]...)
	log.Info(fmt.Sprintf("part 2 solution: %d", result))
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func ParseInput(data []byte) (string, *Route) {
	segments := strings.Split(string(data), "\n\n")
	instructions := segments[0]
	routes := ParseRoutes(segments[1])

	rootRoute, _ := createRoute("AAA", routes, map[string]*Route{})
	return instructions, rootRoute
}

func ParseInputTwo(data []byte) (string, []*Route) {
	segments := strings.Split(string(data), "\n\n")
	instructions := segments[0]

	startingPoints := []string{}
	routes := ParseRoutes(segments[1])
	for name := range routes {
		if (name[len(name)-1]) == 'Z' {
			startingPoints = append(startingPoints, name)
		}
	}

	rootRoutes := []*Route{}
	for _, routeName := range startingPoints {
		route, _ := createRoute(routeName, routes, map[string]*Route{})
		rootRoutes = append(rootRoutes, route)
	}

	return instructions, rootRoutes
}

func createRoute(route string, routes map[string][2]string, knownRoutes map[string]*Route) (*Route, map[string]*Route) {
	r := &Route{
		Name: route,
	}
	leftRoute := routes[route][0]
	rightRoute := routes[route][1]

	knownRoutes[route] = r

	var left *Route
	if leftRoute == route {
		left = nil
	} else {
		left = knownRoutes[leftRoute]
		if left == nil {
			left, knownRoutes = createRoute(leftRoute, routes, knownRoutes)
			knownRoutes[left.Name] = left
		}
	}

	var right *Route
	if rightRoute == route {
		right = nil
	} else {
		right = knownRoutes[rightRoute]
		if right == nil {
			right, knownRoutes = createRoute(rightRoute, routes, knownRoutes)
			knownRoutes[right.Name] = right
		}
	}

	r.Left = left
	r.Right = right

	return r, knownRoutes
}

func ParseRoutes(data string) map[string][2]string {
	out := map[string][2]string{}
	lineRegex := regexp.MustCompile(`(\w+) = \((\w+), (\w+)\)`)
	lines := strings.Split(data, "\n")
	for _, line := range lines {
		matches := lineRegex.FindStringSubmatch(line)
		out[matches[1]] = [2]string{matches[2], matches[3]}
	}
	return out
}
