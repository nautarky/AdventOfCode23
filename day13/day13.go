package day13

func Part1(lines []string) int {
	grids := splitGrids(lines)
	sum := 0
	for _, g := range grids {
		val, isCol := findSymmetry(g)
		if isCol {
			sum += val
		} else {
			sum += 100 * val
		}
	}
	return sum
}

func Part2(lines []string) int {
	return 0
}

// returns (row/col preceding line, isCol?)
func findSymmetry(grid []string) (int, bool) {
	col := findSymmetricColumn(grid)
	if col != -1 {
		return col, true
	}

	return findSymmetricRow(grid), false
}

func findSymmetricColumn(grid []string) int {
	for l := 0; l < len(grid[0])-1; l++ {
		isSymmetric := true

		for _, row := range grid {
			curL, curR := l, l+1

			for isSymmetric && curL >= 0 && curR < len(grid[0]) {
				if row[curL] != row[curR] {
					isSymmetric = false
				}

				curL--
				curR++
			}

			if !isSymmetric {
				break
			}
		}

		if isSymmetric {
			return l + 1
		}
	}

	return -1
}

func findSymmetricRow(grid []string) int {
	for n := 0; n < len(grid)-1; n++ {
		isSymmetric := true

		for col := 0; col < len(grid[0]); col++ {
			curN, curS := n, n+1

			for isSymmetric && curN >= 0 && curS < len(grid) {
				if grid[curN][col] != grid[curS][col] {
					isSymmetric = false
				}

				curN--
				curS++
			}

			if !isSymmetric {
				break
			}
		}

		if isSymmetric {
			return n + 1
		}
	}

	return -1
}

func splitGrids(lines []string) [][]string {
	grids := make([][]string, 0)
	grids = append(grids, make([]string, 0))

	for _, line := range lines {
		if line == "" {
			grids = append(grids, make([]string, 0))
		} else {
			grids[len(grids)-1] = append(grids[len(grids)-1], line)
		}
	}

	return grids
}
