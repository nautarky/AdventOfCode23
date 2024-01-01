package day14

func Part1(lines []string) int {
	grid := buildGrid(lines)
	tiltNorth(grid, len(grid[0]), len(grid))
	return countLoad(grid)
}

func Part2(lines []string) int {
	grid := buildGrid(lines)
	cache := make([][][]byte, 0)

	width, height := len(grid[0]), len(grid)

	iters := 1_000_000_000
	for i := 0; i < iters; i++ {
		cachePos := findGridInCache(grid, cache)
		if cachePos != -1 {
			resultPos := (iters-i)%(len(cache)-cachePos) + cachePos
			return countLoad(cache[resultPos])
		} else {
			dupe := make([][]byte, len(grid))
			for j := range grid {
				dupe[j] = make([]byte, len(grid[j]))
				copy(dupe[j], grid[j])
			}
			cache = append(cache, dupe)
		}
		tiltNorth(grid, width, height)
		tiltWest(grid, width, height)
		tiltSouth(grid, width, height)
		tiltEast(grid, width, height)
	}
	return countLoad(grid)
}

func findGridInCache(grid [][]byte, cache [][][]byte) int {
	for i, cachedGrid := range cache {
		same := true

		for j, row := range cachedGrid {
			for k, val := range row {
				if grid[j][k] != val {
					same = false
					break
				}
			}

			if !same {
				break
			}
		}

		if same {
			return i
		}
	}

	return -1
}

func buildGrid(lines []string) [][]byte {
	grid := make([][]byte, len(lines))

	for i, line := range lines {
		grid[i] = []byte(line)
	}

	return grid
}

func tiltNorth(grid [][]byte, width, height int) {
	for i := 0; i < width; i++ {
		cur := 0

		for cur < height {
			if grid[cur][i] != '.' {
				cur++
				continue
			}

			lead := cur + 1
			for lead < height && grid[lead][i] == '.' {
				lead++
			}

			if lead < height && grid[lead][i] == 'O' {
				grid[cur][i], grid[lead][i] = grid[lead][i], grid[cur][i]
			}

			if lead < height && grid[lead][i] == '#' {
				cur = lead
			}
			cur++
		}
	}
	return
}

func tiltSouth(grid [][]byte, width, height int) {
	for i := 0; i < width; i++ {
		cur := height - 1

		for cur >= 0 {
			if grid[cur][i] != '.' {
				cur--
				continue
			}

			lead := cur - 1
			for lead >= 0 && grid[lead][i] == '.' {
				lead--
			}

			if lead >= 0 && grid[lead][i] == 'O' {
				grid[cur][i], grid[lead][i] = grid[lead][i], grid[cur][i]
			}

			if lead >= 0 && grid[lead][i] == '#' {
				cur = lead
			}
			cur--
		}
	}
	return
}

func tiltEast(grid [][]byte, width, height int) {
	for i := 0; i < height; i++ {
		cur := width - 1

		for cur >= 0 {
			if grid[i][cur] != '.' {
				cur--
				continue
			}

			lead := cur - 1
			for lead >= 0 && grid[i][lead] == '.' {
				lead--
			}

			if lead >= 0 && grid[i][lead] == 'O' {
				grid[i][cur], grid[i][lead] = grid[i][lead], grid[i][cur]
			}

			if lead >= 0 && grid[i][lead] == '#' {
				cur = lead
			}
			cur--
		}
	}
	return
}

func tiltWest(grid [][]byte, width, height int) {

	for i := 0; i < height; i++ {
		cur := 0

		for cur < width {
			if grid[i][cur] != '.' {
				cur++
				continue
			}

			lead := cur + 1
			for lead < width && grid[i][lead] == '.' {
				lead++
			}

			if lead < width && grid[i][lead] == 'O' {
				grid[i][cur], grid[i][lead] = grid[i][lead], grid[i][cur]
			}

			if lead < width && grid[i][lead] == '#' {
				cur = lead
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
