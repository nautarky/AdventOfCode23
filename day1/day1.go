package day1

import (
	"Advent23/shared"
	"strings"
)

var strIntMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
}

type calibration struct {
	value    int
	position int
}

func Part1(lines []string) int {
	sum := 0

	for _, line := range lines {
		first, last := -1, -1

		for _, c := range line {
			val, err := shared.RuneToInt(c)
			if err != nil {
				continue
			}

			if first == -1 {
				first = val
			}

			last = val
		}

		sum += first*10 + last
	}

	return sum
}

func Part2(lines []string) int {
	sum := 0

	for _, line := range lines {
		first := calibration{position: len(line)}
		last := calibration{position: -1}

		for k, v := range strIntMap {
			pos := strings.Index(line, k)

			if pos == -1 {
				continue
			}

			if pos < first.position {
				first = calibration{v, pos}
			}

			pos = strings.LastIndex(line, k)

			if pos > last.position {
				last = calibration{v, pos}
			}
		}

		sum += first.value*10 + last.value
	}

	return sum
}
