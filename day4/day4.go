package day4

import (
	"strings"
)

func Part1(lines []string) int {
	sum := 0

	for _, line := range lines {
		sum += scoreCard(line)
	}

	return sum
}

func Part2(lines []string) int {
	return 0
}

func scoreCard(line string) int {
	parts := strings.Fields(line)
	winningNums := make(map[string]bool, 10)

	i := 2
	for ; parts[i] != "|"; i++ {
		winningNums[parts[i]] = true
	}

	score := 0

	i++ // skip '|'
	for ; i < len(parts); i++ {
		if winningNums[parts[i]] {
			score = max(1, score*2)
		}
	}

	return score
}
