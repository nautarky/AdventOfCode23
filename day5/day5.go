package day5

import (
	"math"
	"strconv"
	"strings"
)

func Part1(lines []string) int {
	seedParts := strings.Fields(lines[0])
	seedRanges := make([]seedRange, 0)

	for _, p := range seedParts {
		start, _ := strconv.Atoi(p)
		seedRanges = append(seedRanges, seedRange{start, 1})
	}

	return solve(lines, seedRanges)
}

func Part2(lines []string) int {
	seedParts := strings.Fields(lines[0])
	seedRanges := make([]seedRange, 0)

	for i := 1; i < len(seedParts); i += 2 {
		start, _ := strconv.Atoi(seedParts[i])
		size, _ := strconv.Atoi(seedParts[i+1])
		seedRanges = append(seedRanges, seedRange{start, size})
	}

	return solve(lines, seedRanges)
}

func solve(lines []string, seedRanges []seedRange) int {
	farmMaps := parseMaps(lines)
	minLoc := math.MaxInt

	for _, sr := range seedRanges {
		for i := sr.start; i < sr.start+sr.size; {
			loc, inc := findLoc(i, farmMaps)
			minLoc = min(minLoc, loc)
			i += inc
		}
	}

	return minLoc
}

type seedRange struct {
	start int
	size  int
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

// returns (location, min_step needed to get to another path)
func findLoc(seed int, maps [][]farmMap) (int, int) {
	minStep := math.MaxInt

	for _, tier := range maps {
		minStepTier := math.MaxInt

		for _, m := range tier {
			if seed >= m.source && seed < m.source+m.size {
				minStepTier = min(minStepTier, m.size-(seed-m.source))
				seed = (seed - m.source) + m.dest
				break
			}
		}

		minStep = min(minStep, minStepTier)
	}

	return seed, minStep
}
