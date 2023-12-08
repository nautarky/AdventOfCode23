package main

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"unicode"
)

const d3InputPath = "input/3.txt"

func d3() (int, int) {
	f, err := os.Open(d3InputPath)
	check(err)
	defer f.Close()

	s := bufio.NewScanner(f)
	schematic := make([][]rune, 0)
	for s.Scan() {
		schematic = append(schematic, []rune(s.Text()))
	}

	solution := &d3Solution{schematic}
	return solution.part1(), solution.part2()
}

type d3Solution struct {
	schematic [][]rune
}

func (s *d3Solution) part1() int {
	sum := 0
	for i := 0; i < len(s.schematic); i++ {
		for j := 0; j < len(s.schematic[i]); j++ {
			if unicode.IsDigit(s.schematic[i][j]) {
				number, leftBound, rightBound, _ := s.findNumberFromPoint(i, j)

				if s.isSymbolAdjacentToRange(i, leftBound, rightBound) {
					sum += number
				}

				j = rightBound
			}
		}
	}

	return sum
}

func (s *d3Solution) part2() int {
	sum := 0

	for i := 0; i < len(s.schematic); i++ {
		for j := 0; j < len(s.schematic[i]); j++ {
			if s.schematic[i][j] == '*' {
				nums := s.findAdjacentNumbers(i, j)
				if len(nums) == 2 {
					sum += nums[0] * nums[1]
				}
			}
		}
	}

	return sum
}

func (s *d3Solution) findAdjacentNumbers(i, j int) []int {
	nums := make([]int, 0)

	for row := max(i-1, 0); row < min(i+2, len(s.schematic)); row++ {
		for col := max(j-1, 0); col < min(j+2, len(s.schematic[i])); col++ {
			if unicode.IsDigit(s.schematic[row][col]) {
				num, _, rightBound, _ := s.findNumberFromPoint(row, col)
				nums = append(nums, num)
				col = rightBound
			}
		}
	}

	return nums
}

// Returns parsed number, left bound, right bound, error msg
func (s *d3Solution) findNumberFromPoint(row, col int) (int, int, int, error) {
	if row < 0 || row >= len(s.schematic) || col < 0 || col >= len(s.schematic[row]) {
		return 0, 0, 0, errors.New("point is out of bounds")
	}

	if !unicode.IsDigit(s.schematic[row][col]) {
		return 0, 0, 0, errors.New("point is not a digit")
	}

	l, r := col, col
	for l > 0 && unicode.IsDigit(s.schematic[row][l-1]) {
		l--
	}

	for r < len(s.schematic[row])-1 && unicode.IsDigit(s.schematic[row][r+1]) {
		r++
	}

	res, err := strconv.Atoi(string(s.schematic[row][l : r+1]))
	return res, l, r, err
}

func (s *d3Solution) isSymbolAdjacentToRange(row, first, last int) bool {
	for i := max(row-1, 0); i < min(row+2, len(s.schematic)); i++ {
		for j := max(first-1, 0); j < min(last+2, len(s.schematic[i])); j++ {
			r := s.schematic[i][j]
			if !unicode.IsDigit(r) && r != '.' {
				return true
			}
		}
	}

	return false
}
