package day03

import "strconv"

func Part1(lines []string) int {
	sum := 0

	for row, line := range lines {
		for col := 0; col < len(line); {
			num, size := getNumberAt(line, col)

			if size > 0 && isSymbolAdjacent(lines, row, col, size) {
				sum += num
			}

			col += size + 1
		}
	}

	return sum
}

func Part2(lines []string) int {
	sum := 0

	for row, line := range lines {
		for col := 0; col < len(line); col++ {
			if line[col] == '*' {
				adjacentNums := getAdjacentNumbers(lines, row, col)

				if len(adjacentNums) == 2 {
					sum += adjacentNums[0] * adjacentNums[1]
				}
			}
		}
	}

	return sum
}

func getAdjacentNumbers(lines []string, row, col int) []int {
	nums := make([]int, 0)

	for i := max(row-1, 0); i < min(row+2, len(lines)); i++ {
		for j := max(col-1, 0); j < min(col+2, len(lines[i])); j++ {
			num, size := getNumberAt(lines[i], j)

			if size > 0 {
				nums = append(nums, num)
			}

			// advance past current number, if one exists
			for j < len(lines[i]) && lines[i][j] >= '0' && lines[i][j] <= '9' {
				j++
			}
		}
	}

	return nums
}

// will grow to the left. returns the value and length. returns (0, 0) for non-numeric positions.
func getNumberAt(line string, pos int) (int, int) {
	l, r := pos, pos

	if line[l] < '0' || line[l] > '9' {
		return 0, 0
	}

	for l > 0 && line[l-1] >= '0' && line[l-1] <= '9' {
		l--
	}

	for r < len(line) && line[r] >= '0' && line[r] <= '9' {
		r++
	}

	num, _ := strconv.Atoi(line[l:r])
	return num, r - l
}

// checks interval [col-1, col+size]
func isSymbolAdjacent(lines []string, row, col, size int) bool {
	for i := max(row-1, 0); i < min(row+2, len(lines)); i++ {
		for j := max(col-1, 0); j < min(col+size+1, len(lines[row])); j++ {
			c := lines[i][j]

			if c != '.' && (c < '0' || c > '9') {
				return true
			}
		}
	}

	return false
}
