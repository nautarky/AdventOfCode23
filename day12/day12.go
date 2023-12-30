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
		sum += rs.countArrangements(0, -1, 0)
	}

	return sum
}

func Part2(lines []string) int {
	sum := 0

	for _, line := range lines {
		rs := newRowSolver(line, 5)
		sum += rs.countArrangements(0, -1, 0)
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
	repeated := make([]string, repeat)
	for i := range repeated {
		repeated[i] = parts[0]
	}
	row := []byte(strings.Join(repeated, "?"))
	spec := findSpec(parts[1], repeat)
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

func (rs *rowSolver) countArrangements(i, groupId, groupLen int) int {
	if i == len(rs.row) {
		if rs.re.Match(rs.row) {
			return 1
		} else {
			return 0
		}
	}

	if groupId >= 0 && (groupId >= len(rs.spec) || groupLen > rs.spec[groupId]) {
		return 0
	}

	if rs.row[i] == '?' {
		sum := 0
		rs.row[i] = '#'
		sum += rs.countArrangements(i, groupId, groupLen)
		rs.row[i] = '.'
		sum += rs.countArrangements(i, groupId, groupLen)
		rs.row[i] = '?'
		return sum
	}

	if rs.row[i] == '#' {
		if i == 0 {
			groupId++
			groupLen++
		} else if rs.row[i-1] == '#' {
			groupLen++
		} else if rs.row[i-1] == '.' {
			groupId++
			groupLen = 1
		}
	} else {
		groupLen = 0
	}

	return rs.countArrangements(i+1, groupId, groupLen)
}

func findSpec(line string, repeat int) []int {
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
