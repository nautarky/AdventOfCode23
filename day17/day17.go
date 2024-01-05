package day17

import (
	"math"
	"slices"
)

func Part1(lines []string) int {
	sol := newSolution(lines)
	sol.dfs(point{1, 0}, point{0, 0}, 0, 1)
	sol.dfs(point{0, 1}, point{0, 0}, 0, 1)
	// shared.PrintNestedSliceInt(sol.cost)
	// shared.PrintNestedSliceInt(sol.minCost)
	return sol.minCost[len(sol.minCost)-1][len(sol.minCost[0])-1]
}

func Part2(lines []string) int {
	return 0
}

type solution struct {
	cost    [][]int
	minCost [][]int
}

func newSolution(lines []string) solution {
	cost := make([][]int, len(lines))
	for i, line := range lines {
		row := make([]int, len(line))
		for j, val := range line {
			row[j] = int(val - '0')
		}
		cost[i] = row
	}

	minCost := make([][]int, len(lines))
	for i, line := range lines {
		row := make([]int, len(line))
		for j := range line {
			row[j] = math.MaxInt
		}
		minCost[i] = row
	}

	minCost[0][0] = 0
	return solution{cost, minCost}
}

func (s *solution) dfs(cur, prev point, runCost, streak int) {
	curCost := runCost + s.cost[cur.row][cur.col]
	if streak > 3 || curCost > s.minCost[cur.row][cur.col] {
		return
	}
	s.minCost[cur.row][cur.col] = curCost

	points := s.nextPoints(cur, prev)
	for _, p := range points {
		nextStreak := 1
		if isStreak(p, prev) {
			nextStreak = streak + 1
		}
		s.dfs(p, cur, curCost, nextStreak)
	}
}

func (s *solution) nextPoints(cur, prev point) []point {
	points := []point{{cur.row, cur.col - 1}, {cur.row, cur.col + 1}, {cur.row - 1, cur.col}, {cur.row + 1, cur.col}}

	return slices.DeleteFunc(points, func(p point) bool {
		return p.col < 0 || p.col >= len(s.minCost[0]) || p.row < 0 || p.row >= len(s.minCost) || p == prev
	})
}

func isStreak(next, prev point) bool {
	return next.row == prev.row+2 || next.row == prev.row-2 || next.col == prev.col+2 || next.col == prev.col-2
}

type point struct {
	row int
	col int
}
