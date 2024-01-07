package day08

import (
	"Advent23/shared"
	"strings"
)

func Part1(lines []string) int {
	graph := buildGraph(lines[2:])
	cur := "AAA"
	stop := func(s string) bool { return s == "ZZZ" }
	return findCycleLen(graph, cur, lines[0], stop)
}

func Part2(lines []string) int {
	graph := buildGraph(lines[2:])
	positions := findStarts(graph)
	cycles := make([]int, len(positions))
	stop := func(s string) bool { return s[len(s)-1] == 'Z' }

	for i, pos := range positions {
		cycles[i] = findCycleLen(graph, pos, lines[0], stop)
	}

	return shared.LCM(cycles[0], cycles[1], cycles[2:]...)
}

func buildGraph(lines []string) map[string][2]string {
	output := make(map[string][2]string, len(lines))

	for _, line := range lines {
		parts := strings.Fields(line)
		output[parts[0]] = [2]string{parts[2][1:4], parts[3][0:3]}
	}

	return output
}

func findCycleLen(graph map[string][2]string, start string, instructions string, stop func(string) bool) int {
	cur, steps := start, 0

	for {
		instruction := instructions[steps%len(instructions)]

		var place int
		if instruction == 'R' {
			place = 1
		}

		cur = graph[cur][place]
		steps++

		if stop(cur) {
			return steps
		}
	}
}

func findStarts(graph map[string][2]string) []string {
	starts := make([]string, 0)

	for k := range graph {
		if k[len(k)-1] == 'A' {
			starts = append(starts, k)
		}
	}

	return starts
}
