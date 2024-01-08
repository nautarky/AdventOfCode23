package day18

import (
	"Advent23/shared"
	"fmt"
	"strconv"
	"strings"
)

func Part1(lines []string) int {
	instructions := parseInstructions1(lines)
	return solve(instructions)
}

func Part2(lines []string) int {
	instructions := parseInstructions2(lines)
	return solve(instructions)
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
	picks := shared.PicksTheoremIInt(area, len(points))
	fmt.Println(len(points))
	return len(points) + int(picks)
}

func parseInstructions1(lines []string) []instruction {
	instructions := make([]instruction, len(lines))
	for i, line := range lines {
		parts := strings.Fields(line)
		direction := dirToComplex(parts[0], 0)
		length, _ := strconv.Atoi(parts[1])
		instructions[i] = instruction{direction, length}
	}
	return instructions
}

func parseInstructions2(lines []string) []instruction {
	instructions := make([]instruction, len(lines))
	for i, line := range lines {
		parts := strings.Fields(line)
		ins := parts[2][2 : len(parts[2])-1]
		dirI, _ := strconv.Atoi(ins[5:])
		dir := dirToComplex("", dirI)
		length, _ := strconv.ParseInt(ins[:5], 16, 64)
		instructions[i] = instruction{dir, int(length)}
	}
	return instructions
}

func dirToComplex(dirS string, dirI int) complex64 {
	switch dirS {
	case "R":
		return complex(1, 0)
	case "D":
		return complex(0, -1)
	case "L":
		return complex(-1, 0)
	case "U":
		return complex(0, 1)
	}

	switch dirI {
	case 0:
		return complex(1, 0)
	case 1:
		return complex(0, -1)
	case 2:
		return complex(-1, 0)
	default:
		return complex(0, 1)
	}
}

type instruction struct {
	direction complex64
	length    int
}
