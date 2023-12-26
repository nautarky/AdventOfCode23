package day4

import (
	"math"
	"strings"
)

func Part1(lines []string) int {
	sum := 0

	for _, line := range lines {
		sum += score(countMatches(line))
	}

	return sum
}

func Part2(lines []string) int {
	copies := make([]int, len(lines))
	for i := 0; i < len(copies); i++ {
		copies[i] = 1
	}

	for i, line := range lines {
		matches := countMatches(line)

		for j := i + 1; j-i <= matches; j++ {
			copies[j] += copies[i]
		}
	}

	sum := 0
	for _, count := range copies {
		sum += count
	}

	return sum
}

func countMatches(line string) int {
	parts := strings.Fields(line)
	winningNums := make(map[string]bool, 10)

	i := 2
	for ; parts[i] != "|"; i++ {
		winningNums[parts[i]] = true
	}

	matches := 0

	i++ // skip '|'
	for ; i < len(parts); i++ {
		if winningNums[parts[i]] {
			matches++
		}
	}

	return matches
}

func score(matches int) int {
	if matches == 0 {
		return 0
	}

	return int(math.Exp2(float64(matches - 1)))
}
