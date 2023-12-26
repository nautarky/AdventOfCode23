package day2

import (
	"strconv"
	"strings"
)

func Part1(lines []string) int {
	sum := 0

	for i, line := range lines {
		g := buildGame(line)
		if g["red"] <= 12 && g["blue"] <= 14 && g["green"] <= 13 {
			sum += i + 1
		}
	}

	return sum
}

func Part2(lines []string) int {
	return 0
}

// stores a game's minimum cube count by color
type game map[string]int

func buildGame(line string) game {
	g := make(game, 3)

	parts := strings.FieldsFunc(line, split)
	for i := 2; i < len(parts); i += 2 {
		num, color := parts[i], parts[i+1]
		count, _ := strconv.Atoi(num)
		g[color] = max(g[color], count)
	}

	return g
}

func split(r rune) bool {
	return r == ':' || r == ';' || r == ',' || r == ' '
}
