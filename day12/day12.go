package day12

import (
	"Advent23/shared"
	"slices"
	"strconv"
	"strings"
)

func Part1(lines []string) int {
	sum := 0

	for _, line := range lines {
		rs := newRowSolver(line, 1)
		sum += rs.countArrangements(0, 0, 0, '.')
	}

	return sum
}

func Part2(lines []string) int {
	return 0
}

type rowSolver struct {
	row    []byte
	groups []int
}

func newRowSolver(line string, repeat int) *rowSolver {
	parts := strings.Fields(line)
	repeated := make([]string, repeat)
	for i := range repeated {
		repeated[i] = parts[0]
	}
	row := []byte(strings.Join(repeated, "?"))
	row = append(row, '.')
	groups := parseGroups(parts[1], repeat)
	rs := rowSolver{row: row, groups: groups}
	return &rs
}

func (rs *rowSolver) countArrangements(i, group, run int, prev byte) int {
	// groups overrunneth
	if group >= len(rs.groups) || run > rs.groups[group] {
		return 0
	}
	// not enough string left
	if len(rs.row)-i < shared.SumIntSlice(rs.groups[group:])-run {
		return 0
	}

	// all groups are satisfied by s[:i]
	if group == len(rs.groups)-1 && run == rs.groups[len(rs.groups)-1] {
		if slices.Contains(rs.row[i:], '#') {
			return 0
		} else {
			return 1
		}
	}

	b := rs.row[i]
	if b == '.' && prev == '#' && run != rs.groups[group] {
		return 0
	} else if b == '.' && prev == '#' {
		return rs.countArrangements(i+1, group+1, 0, '.')
	} else if b == '.' && prev == '.' {
		return rs.countArrangements(i+1, group, 0, '.')
	} else if b == '#' {
		return rs.countArrangements(i+1, group, run+1, '#')
	}

	// b == '?', so do both
	sum := 0
	rs.row[i] = '.'
	sum += rs.countArrangements(i, group, run, prev)
	rs.row[i] = '#'
	sum += rs.countArrangements(i, group, run, prev)
	rs.row[i] = '?'
	return sum
}

func parseGroups(line string, repeat int) []int {
	nums := strings.Split(line, ",")

	output := make([]int, 0)

	for i := 0; i < repeat; i++ {
		for _, num := range nums {
			val, _ := strconv.Atoi(num)
			output = append(output, val)
		}
	}

	return output
}
