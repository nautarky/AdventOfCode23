package day12

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func Part1(lines []string) int {
	sum := 0

	for _, line := range lines {
		parts := strings.Fields(line)
		spec := findSpec(parts[1])
		re := buildRegex(spec)
		re.MatchString(parts[0])
		sum += countArrangements([]byte(parts[0]), 0, re)
	}

	return sum
}

func Part2(lines []string) int {
	return 0
}

func buildRegex(spec []int) *regexp.Regexp {
	var sb strings.Builder

	sb.WriteString(`^\.*`)
	for i, num := range spec {
		sb.WriteString(fmt.Sprintf("#{%d}", num))
		if i != len(spec)-1 {
			sb.WriteString(`\.+`)
		}
	}
	sb.WriteString(`\.*$`)

	return regexp.MustCompile(sb.String())
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

func countArrangements(row []byte, start int, re *regexp.Regexp) int {
	if !slices.Contains(row, '?') {
		if re.Match(row) {
			return 1
		} else {
			return 0
		}
	}

	sum := 0

	for i := start; i < len(row); i++ {
		if row[i] == '?' {
			row[i] = '#'
			sum += countArrangements(row, i+1, re)
			row[i] = '.'
			sum += countArrangements(row, i+1, re)
			row[i] = '?'
		}
	}

	return sum
}
