package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type d5 struct {
	seeds []int64
	maps  [][][3]int64
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

	maps := make([][][3]int64, 0)
	for s.Scan() {
		val := s.Text()

		// Title line
		if val == "" {
			maps = append(maps, make([][3]int64, 0))
			s.Scan()
			continue
		}

		rawMap := strings.Fields(val)
		dest, _ := strconv.ParseInt(rawMap[0], 10, 64)
		source, _ := strconv.ParseInt(rawMap[1], 10, 64)
		size, _ := strconv.ParseInt(rawMap[2], 10, 64)
		maps[len(maps)-1] = append(maps[len(maps)-1], [3]int64{dest, source, size})
	}

	err = file.Close()
	check(err)

	return &d5{seeds, maps}
}

func (d *d5) part1() int64 {
	locations := make([]int64, 0)

	for _, seed := range d.seeds {
		for _, m := range d.maps {
			seed = mapValue(seed, m)
		}

		locations = append(locations, seed)
	}

	lowest := locations[0]
	for _, l := range locations[1:] {
		lowest = min(lowest, l)
	}

	return lowest
}

func mapValue(value int64, m [][3]int64) int64 {
	for _, r := range m {
		if value >= r[1] && value <= r[1]+r[2] {
			offset := value - r[1]
			return r[0] + offset
		}
	}

	return value
}
