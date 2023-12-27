package day5

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

func Part1(lines []string) int {
	seedParts := strings.Fields(lines[0])
	seeds := make([]int, 0)

	for _, s := range seedParts[1:] {
		seed, _ := strconv.Atoi(s)
		seeds = append(seeds, seed)
	}

	slices.Sort(seeds)
	farmMaps := parseMaps(lines)
	breaks := findBreaks(farmMaps)

	minLoc, curBreak := math.MaxInt, 0
	for _, seed := range seeds {
		if seed < breaks[curBreak] {
			continue
		}

		// process this seed
		minLoc = min(minLoc, findLoc(seed, farmMaps))

		for seed > breaks[curBreak] {
			curBreak++
		}
	}

	return minLoc
}

func Part2(lines []string) int {
	return 0
}

type farmMap struct {
	source int
	dest   int
	size   int
}

func parseMaps(lines []string) [][]farmMap {
	farmMaps := make([][]farmMap, 0)

	for _, line := range lines[1:] {
		if line == "" {
			continue
		}

		if strings.Contains(line, "map:") {
			farmMaps = append(farmMaps, make([]farmMap, 0))
			continue
		}

		farmMaps[len(farmMaps)-1] = append(farmMaps[len(farmMaps)-1], parseMap(line))
	}

	return farmMaps
}

func parseMap(line string) farmMap {
	parts := strings.Fields(line)

	dest, _ := strconv.Atoi(parts[0])
	source, _ := strconv.Atoi(parts[1])
	size, _ := strconv.Atoi(parts[2])

	return farmMap{source, dest, size}
}

// returns sorted slice of domain interval openings that lead to different ranges, starting at location and
// moving backward.
func findBreaks(maps [][]farmMap) []int {
	breaks := map[int]bool{
		0:           true,
		math.MaxInt: true,
	}

	for i := len(maps) - 1; i >= 0; i-- {
		for j := 0; j < len(maps[i]); j++ {
			m := maps[i][j]

			if !breaks[m.source] {
				breaks[m.source] = true
			}

			if !breaks[m.source+m.size] {
				breaks[m.source+m.size] = true
			}
		}
	}

	output := make([]int, 0)
	for k := range breaks {
		output = append(output, k)
	}

	slices.Sort(output)
	return output
}

func findLoc(seed int, maps [][]farmMap) int {
	for _, tier := range maps {
		for _, m := range tier {
			if seed > m.source && seed < m.source+m.size {
				seed = (seed - m.source) + m.dest
				break
			}
		}
	}

	return seed
}
