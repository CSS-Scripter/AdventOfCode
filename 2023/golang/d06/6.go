package d06

import (
	"fmt"
	"math"

	"github.com/charmbracelet/log"
)

type Race struct {
	Duration float64
	Distance float64
}

func Main() {
	one([]Race{
		{47, 282},
		{70, 1079},
		{75, 1147},
		{66, 1062},
	})

	two(Race{47707566, 282107911471062})
}

func one(races []Race) {
	var sum float64 = 1
	for _, race := range races {
		first, second := solveABC(race)
		tolerance := second - first + 1
		sum *= tolerance
	}
	log.Info(fmt.Sprintf("part 1 solution: %d", int(sum)))
}

func two(race Race) {
	first, second := solveABC(race)
	tolerance := second - first + 1
	log.Info(fmt.Sprintf("part 2 solution: %d", int(tolerance)))
}

func solveABC(race Race) (float64, float64) {
	var a float64 = 1
	b := -race.Duration
	c := race.Distance + 1
	d := (b * b) - (4 * a * c)

	sqrtD := math.Sqrt(d)

	first := (-b + sqrtD) / (2 * a)
	second := (-b - sqrtD) / (2 * a)

	return math.Ceil(math.Min(first, second)), math.Floor(math.Max(first, second))
}
