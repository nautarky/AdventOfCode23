package day12

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Part1(lines []string) int {
	sum := 0

	for _, line := range lines {
		rs := newRowSolver(line, 1)
		sum += rs.countArrangements(0)
	}

	return sum
}

func Part2(lines []string) int {
	sum := 0

	for _, line := range lines {
		rs := newRowSolver(line, 5)
		sum += rs.countArrangements(0)
	}

	return sum
}

type rowSolver struct {
	row  []byte
	spec []int
	re   *regexp.Regexp
}

func newRowSolver(line string, repeat int) *rowSolver {
	parts := strings.Fields(line)
	row := []byte(strings.Repeat(parts[0], repeat))
	spec := findSpec(strings.Repeat(parts[1], repeat))
	rs := rowSolver{row: row, spec: spec}
	rs.buildRegex()
	return &rs
}

func (rs *rowSolver) buildRegex() {
	var sb strings.Builder

	sb.WriteString(`^\.*`)
	for i, num := range rs.spec {
		sb.WriteString(fmt.Sprintf("#{%d}", num))
		if i != len(rs.spec)-1 {
			sb.WriteString(`\.+`)
		}
	}
	sb.WriteString(`\.*$`)

	rs.re = regexp.MustCompile(sb.String())
}

func (rs *rowSolver) countArrangements(i int) int {
	if i == len(rs.row) {
		if rs.re.Match(rs.row) {
			return 1
		} else {
			return 0
		}
	}

	if rs.row[i] != '?' {
		return rs.countArrangements(i + 1)
	}

	sum := 0
	rs.row[i] = '#'
	sum += rs.countArrangements(i + 1)
	rs.row[i] = '.'
	sum += rs.countArrangements(i + 1)
	rs.row[i] = '?'

	return sum
}

func findSpec(line string) []int {
	nums := strings.Split(line, ",")

	output := make([]int, len(nums))

	for i, num := range nums {
		val, _ := strconv.Atoi(num)
		output[i] = val
	}

	return output
}
