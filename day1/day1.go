package day1

import "Advent23/shared"

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
