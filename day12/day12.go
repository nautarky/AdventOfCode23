package day12

import (
	"Advent23/shared"
	"fmt"
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
	sum := 0

	for _, line := range lines {
		rs := newRowSolver(line, 5)
		sum += rs.countArrangements(0, 0, 0, '.')
	}

	return sum
}

type rowSolver struct {
	row    []byte
	groups []int
	memo   map[string]int
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
	memo := make(map[string]int)
	rs := rowSolver{row: row, groups: groups, memo: memo}
	return &rs
}

func (rs *rowSolver) countArrangements(i, group, run int, prev byte) int {
	if i >= len(rs.row) {
		return 0
	}

	key := fmt.Sprintf("%d|%d|%d|%c|%c", i, group, run, prev, rs.row[i])
	cache, ok := rs.memo[key]
	if ok {
		return cache
	}

	// groups overrunneth
	if group >= len(rs.groups) || run > rs.groups[group] {
		rs.memo[key] = 0
		return 0
	}
	// not enough string left
	if len(rs.row)-i < shared.SumIntSlice(rs.groups[group:])-run {
		rs.memo[key] = 0
		return 0
	}

	// all groups are satisfied by s[:i]
	if group == len(rs.groups)-1 && run == rs.groups[len(rs.groups)-1] {
		if slices.Contains(rs.row[i:], '#') {
			rs.memo[key] = 0
			return 0
		} else {
			rs.memo[key] = 1
			return 1
		}
	}

	b := rs.row[i]
	if b == '.' && prev == '#' && run != rs.groups[group] {
		return 0
	} else if b == '.' && prev == '#' {
		res := rs.countArrangements(i+1, group+1, 0, '.')
		rs.memo[key] = res
		return res
	} else if b == '.' && prev == '.' {
		res := rs.countArrangements(i+1, group, 0, '.')
		rs.memo[key] = res
		return res
	} else if b == '#' {
		res := rs.countArrangements(i+1, group, run+1, '#')
		rs.memo[key] = res
		return res
	}

	// b == '?', so do both
	res := 0
	rs.row[i] = '.'
	res += rs.countArrangements(i, group, run, prev)
	rs.row[i] = '#'
	res += rs.countArrangements(i, group, run, prev)
	rs.row[i] = '?'
	rs.memo[key] = res
	return res
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
