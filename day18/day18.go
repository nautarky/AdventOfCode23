package day18

import (
	"Advent23/shared"
	"strconv"
	"strings"
)

func Part1(lines []string) int {
	instructions := parseInstructions1(lines)
	return solve(instructions)
}

func Part2(lines []string) int {
	return 0
}

func solve(instructions []instruction) int {
	points := []complex64{}

	var cur complex64 = complex(0, 0)
	for _, ins := range instructions {
		for i := 0; i < ins.length; i++ {
			cur += ins.direction
			points = append(points, cur)
		}
	}

	area := shared.ShoelaceComplex(points)
	picks := shared.PicksTheoremI(area, float64(len(points)))
	return len(points) + int(picks)
}

func parseInstructions1(lines []string) []instruction {
	instructions := make([]instruction, len(lines))
	for i, line := range lines {
		parts := strings.Fields(line)
		direction := dirToComplex(parts[0])
		length, _ := strconv.Atoi(parts[1])
		instructions[i] = instruction{direction, length}
	}
	return instructions
}

func dirToComplex(dir string) complex64 {
	switch dir {
	case "R":
		return complex(1, 0)
	case "D":
		return complex(0, -1)
	case "L":
		return complex(-1, 0)
	default:
		return complex(0, 1)
	}
}

type instruction struct {
	direction complex64
	length    int
}
