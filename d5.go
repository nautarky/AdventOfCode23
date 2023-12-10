package main

import (
	"bufio"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type almanacRange struct {
	src  int64
	dest int64
	size int64
}

type d5 struct {
	seeds []int64
	maps  [][]almanacRange
}

func newD5(path string) *d5 {
	file, err := os.Open(path)
	check(err)

	s := bufio.NewScanner(file)
	s.Scan()

	rawSeeds := strings.Fields(s.Text())
	seeds := make([]int64, 0)
	for _, s := range rawSeeds[1:] {
		id, _ := strconv.ParseInt(s, 10, 64)
		seeds = append(seeds, id)
	}

	maps := make([][]almanacRange, 0)

	for s.Scan() {
		val := s.Text()

		// Title line
		if val == "" {
			maps = append(maps, make([]almanacRange, 0))
			s.Scan()
			continue
		}

		rawMap := strings.Fields(val)
		dest, _ := strconv.ParseInt(rawMap[0], 10, 64)
		source, _ := strconv.ParseInt(rawMap[1], 10, 64)
		size, _ := strconv.ParseInt(rawMap[2], 10, 64)
		maps[len(maps)-1] = append(maps[len(maps)-1], almanacRange{source, dest, size})
	}

	err = file.Close()
	check(err)

	for _, m := range maps {
		sort.Sort(ByDestination(m))
	}

	return &d5{seeds, maps}
}

type ByDestination []almanacRange

func (d ByDestination) Len() int           { return len(d) }
func (d ByDestination) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }
func (d ByDestination) Less(i, j int) bool { return d[i].dest < d[j].dest }

func (d *d5) part1() int64 {
	var lowest int64 = math.MaxInt64

	for _, seed := range d.seeds {
		for _, m := range d.maps {
			seed = lookup(seed, m)
		}

		lowest = min(lowest, seed)
	}

	return lowest
}

func (d *d5) part2() int64 {
	locationLevel := len(d.maps) - 1

	for i := int64(1); i < d.maps[locationLevel][0].dest; i++ {
		if d.leadsToSeed(i, locationLevel) {
			return i
		}
	}

	for _, r := range d.maps[locationLevel] {
		for i := r.dest; i < r.dest+r.size; i++ {
			if d.leadsToSeed(i, locationLevel) {
				return i
			}
		}
	}

	return -1
}

// level: current position in map array
func (d *d5) leadsToSeed(dest int64, level int) bool {
	src := dest

	// find range at this level that contains dest
	for _, r := range d.maps[level] {
		if isDestInRange(dest, r) {
			src = reverseLookup(dest, r)
		}
	}

	// exit: does this range lead to a seed?
	if level == 0 {
		return d.isValidSeed(src)
	}

	return d.leadsToSeed(src, level-1)
}

func (d *d5) isValidSeed(seed int64) bool {
	for i := 0; i < len(d.seeds); i += 2 {
		if seed >= d.seeds[i] && seed < d.seeds[i]+d.seeds[i+1] {
			return true
		}
	}

	return false
}

func lookup(src int64, m []almanacRange) int64 {
	for _, r := range m {
		if src >= r.src && src <= r.src+r.size {
			offset := src - r.src
			return r.dest + offset
		}
	}

	return src
}

func reverseLookup(dest int64, r almanacRange) int64 {
	if isDestInRange(dest, r) {
		offset := dest - r.dest
		return r.src + offset
	}

	return dest
}

func isDestInRange(dest int64, r almanacRange) bool {
	return dest >= r.dest && dest <= r.dest+r.size
}
