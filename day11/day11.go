package day11

import (
	"Advent23/shared"
	"math"
	"strings"
)

func Part1(lines []string) int {
	cols := findExpandingCols(lines)
	rows := findExpandingRows(lines)
	points := findPoints(lines, cols, rows)

	sum := 0

	for i, p1 := range points[:len(points)-1] {
		for _, p2 := range points[i+1:] {
			sum += int(math.Abs(float64(p1.X-p2.X)) + math.Abs(float64(p1.Y-p2.Y)))
		}
	}

	return sum
}

func Part2(lines []string) int {
	return 0
}

func findExpandingRows(lines []string) map[int]bool {
	rows := make(map[int]bool)

	for i, row := range lines {
		if strings.Contains(row, "#") {
			continue
		}
		rows[i] = true
	}

	return rows
}

func findExpandingCols(lines []string) map[int]bool {
	cols := make(map[int]bool)

	for i := 0; i < len(lines[0]); i++ {
		for j := 0; j < len(lines); j++ {
			if lines[j][i] == '#' {
				break
			}

			if j == len(lines)-1 {
				cols[i] = true
			}
		}
	}

	return cols
}

func findPoints(lines []string, bigCols map[int]bool, bigRows map[int]bool) []shared.Point {
	points := make([]shared.Point, 0)

	for y, grownY := 0, 0; y < len(lines); y++ {
		if bigRows[y] {
			grownY++
			continue
		}

		for x, grownX := 0, 0; x < len(lines[y]); x++ {
			if bigCols[x] {
				grownX++
				continue
			}

			if lines[y][x] == '#' {
				points = append(points, shared.Point{X: x + grownX, Y: y + grownY})
			}
		}
	}

	return points
}
