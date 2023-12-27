package day8

import "strings"

func Part1(lines []string) int {
	graph := buildGraph(lines[2:])
	cur, steps := "AAA", 0

	for cur != "ZZZ" {
		instruction := lines[0][steps%len(lines[0])]

		if instruction == 'L' {
			cur = graph[cur][0]
		} else {
			cur = graph[cur][1]
		}

		steps++
	}

	return steps
}

func Part2(lines []string) int {
	return 0
}

func buildGraph(lines []string) map[string][2]string {
	output := make(map[string][2]string, len(lines))

	for _, line := range lines {
		parts := strings.Fields(line)
		output[parts[0]] = [2]string{parts[2][1:4], parts[3][0:3]}
	}

	return output
}
