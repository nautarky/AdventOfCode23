package day14

func Part1(lines []string) int {
	grid := buildGrid(lines)
	tiltNorth(grid)
	return countLoad(grid)
}

func Part2(lines []string) int {
	return 0
}

func buildGrid(lines []string) [][]byte {
	grid := make([][]byte, len(lines))

	for i, line := range lines {
		grid[i] = []byte(line)
	}

	return grid
}

func tiltNorth(grid [][]byte) {
	for i := 0; i < len(grid[0]); i++ {
		cur := 0

		for cur < len(grid) {
			if grid[cur][i] != '.' {
				cur++
				continue
			}

			lead := cur + 1
			for lead < len(grid) && grid[lead][i] == '.' {
				lead++
			}

			if lead < len(grid) && grid[lead][i] == 'O' {
				grid[cur][i], grid[lead][i] = grid[lead][i], grid[cur][i]
			}
			cur++
		}
	}
	return
}

func countLoad(grid [][]byte) int {
	sum := 0

	for i, row := range grid {
		for _, val := range row {
			if val == 'O' {
				sum += len(grid) - i
			}
		}
	}

	return sum
}
