package day13

func Part1(lines []string) int {
	return solve(lines, 0)
}

func Part2(lines []string) int {
	return solve(lines, 1)
}

func solve(lines []string, differences int) int {
	grids := splitGrids(lines)
	sum := 0
	for _, g := range grids {
		val, isCol := findSymmetry(g, differences)
		if isCol {
			sum += val
		} else {
			sum += 100 * val
		}
	}
	return sum
}

// returns (row/col preceding line, isCol?)
func findSymmetry(grid []string, differences int) (int, bool) {
	col := findSymmetricColumn(grid, differences)
	if col != -1 {
		return col, true
	}

	return findSymmetricRow(grid, differences), false
}

func findSymmetricColumn(grid []string, differences int) int {
	for l := 0; l < len(grid[0])-1; l++ {
		isSymmetric := true
		curDiff := differences

		for _, row := range grid {
			curL, curR := l, l+1

			for isSymmetric && curL >= 0 && curR < len(grid[0]) {
				if row[curL] != row[curR] {
					if curDiff == 0 {
						isSymmetric = false
					} else {
						curDiff--
					}
				}

				curL--
				curR++
			}

			if !isSymmetric {
				break
			}
		}

		if isSymmetric && curDiff == 0 {
			return l + 1
		}
	}

	return -1
}

func findSymmetricRow(grid []string, differences int) int {
	for n := 0; n < len(grid)-1; n++ {
		isSymmetric := true
		curDiff := differences

		for col := 0; col < len(grid[0]); col++ {
			curN, curS := n, n+1

			for isSymmetric && curN >= 0 && curS < len(grid) {
				if grid[curN][col] != grid[curS][col] {
					if curDiff == 0 {
						isSymmetric = false
					} else {
						curDiff--
					}
				}

				curN--
				curS++
			}

			if !isSymmetric {
				break
			}
		}

		if isSymmetric && curDiff == 0 {
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
