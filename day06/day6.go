package day06

import (
	"Advent23/shared"
	"math"
	"strconv"
	"strings"
)

func Part1(lines []string) int {
	timeParts := strings.Fields(lines[0])
	distParts := strings.Fields(lines[1])

	races := make([]race, len(timeParts)-1)
	for i := 1; i < len(timeParts); i++ {
		t, _ := strconv.Atoi(timeParts[i])
		d, _ := strconv.Atoi(distParts[i])
		races[i-1] = race{t, d}
	}

	product := 1

	for _, r := range races {
		product *= minWays(r)
	}

	return product
}

func Part2(lines []string) int {
	timeParts := strings.Fields(lines[0])
	distParts := strings.Fields(lines[1])

	t, _ := strconv.Atoi(strings.Join(timeParts[1:], ""))
	d, _ := strconv.Atoi(strings.Join(distParts[1:], ""))

	return minWays(race{t, d})
}

type race struct {
	time     int
	distance int
}

func minWays(r race) int {
	a, b := shared.QuadraticFormula(-1.0, float64(r.time), float64(-r.distance))

	lo := math.Ceil(min(a, b) + 0.0000000001)
	hi := math.Floor(max(a, b) - 0.0000000001)

	return int(hi) - int(lo) + 1
}
