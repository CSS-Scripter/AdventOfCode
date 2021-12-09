package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

var (
	requirements  []Requirement
	nearbyTickets []Ticket
)

type Requirement struct {
	name   string
	values []RequirementMinMax
}

type RequirementMinMax struct {
	min int
	max int
}

type Ticket struct {
	values []int
}

func (r Requirement) isValueValid(value int) bool {
	for _, minmax := range r.values {
		if value >= minmax.min && value <= minmax.max {
			return true
		}
	}
	return false
}

func (t Ticket) doesTicketMatchRequirement(r Requirement) (bool, []int) {
	invalidValues := []int{}
	for _, ticketValue := range t.values {
		if !r.isValueValid(ticketValue) {
			invalidValues = append(invalidValues, ticketValue)
		}
	}
	return len(invalidValues) > 0, invalidValues
}

func main() {
	one()
}

func one() {
	prepareInput("example_input.txt")
	for _, ticket := range nearbyTickets {
		invalidTicketValues := []int{}
		for _, req := range requirements {
			invalid, values := ticket.doesTicketMatchRequirement(req)
			if invalid {
				invalidTicketValues = append(invalidTicketValues, values...)
			}
		}
	}
}

func prepareInput(file string) {
	input, _ := ioutil.ReadFile(file)
	lines := strings.Split(strings.TrimSpace(string(input)), "\r\n\r\n")
	parseRequirements(lines[0])
}

func parseRequirements(requirementLines string) {
	for _, requirementLine := range strings.Split(requirementLines, "\r\n") {
		name := strings.Split(requirementLine, ": ")[0]
		values := strings.Split(requirementLine, ": ")[1]
		requirement := Requirement{name: name, values: []RequirementMinMax{}}
		for _, minmax := range strings.Split(values, " or ") {
			min, _ := strconv.Atoi(strings.Split(minmax, "-")[0])
			max, _ := strconv.Atoi(strings.Split(minmax, "-")[1])
			requirement.values = append(requirement.values, RequirementMinMax{min, max})
		}
		requirements = append(requirements, requirement)
	}
}

func parseNearbyTickets(ticketsLine string) {
	for _, ticketLine := range strings.Split(ticketsLine, "\r\n")[1:] {
		ticket := Ticket{}
		for _, value := range strings.Split(ticketLine, ",") {
			valInt, _ := strconv.Atoi(value)
			ticket.values = append(ticket.values, valInt)
		}
		nearbyTickets = append(nearbyTickets, ticket)
	}
}
