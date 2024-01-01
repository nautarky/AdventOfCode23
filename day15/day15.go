package day15

func Part1(lines []string) int {
	sum, cur := 0, 0

	for _, b := range lines[0] {
		if b == ',' {
			sum += cur
			cur = 0
			continue
		}

		cur = ((cur + int(b)) * 17) % 256
	}

	return sum + cur
}

func Part2(lines []string) int {
	return 0
}
