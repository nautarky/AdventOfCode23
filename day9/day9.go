package day9

import (
	"strconv"
	"strings"
)

func Part1(lines []string) int {
	return solve(lines, extrapolateForward)
}

func Part2(lines []string) int {
	return solve(lines, extrapolateBackward)
}

func solve(lines []string, extrapolate func([]int) int) int {
	slices := buildIntSlices(lines)
	sum := 0

	for _, s := range slices {
		sum += extrapolate(s)
	}

	return sum
}

func buildIntSlices(lines []string) [][]int {
	slices := make([][]int, len(lines))

	for i, line := range lines {
		parts := strings.Fields(line)
		s := make([]int, len(parts))

		for j, p := range parts {
			val, _ := strconv.Atoi(p)
			s[j] = val
		}

		slices[i] = s
	}

	return slices
}

func extrapolateBackward(s []int) int {
	if allZeroes(s) {
		return 0
	}

	next := make([]int, len(s)-1)
	for i := 1; i < len(s); i++ {
		next[i-1] = s[i] - s[i-1]
	}

	return s[0] - extrapolateBackward(next)
}

func extrapolateForward(s []int) int {
	if allZeroes(s) {
		return 0
	}

	next := make([]int, len(s)-1)
	for i := 1; i < len(s); i++ {
		next[i-1] = s[i] - s[i-1]
	}

	return s[len(s)-1] + extrapolateForward(next)
}

func allZeroes(s []int) bool {
	for _, i := range s {
		if i != 0 {
			return false
		}
	}
	return true
}
