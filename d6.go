package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type race struct {
	time int
	dist int
}

type d6 struct {
	p1Races []race
	p2Race  race
}

func newD6(path string) *d6 {
	file, err := os.Open(path)
	check(err)

	s := bufio.NewScanner(file)
	s.Scan()
	times := strings.Fields(s.Text())
	s.Scan()
	distances := strings.Fields(s.Text())
	file.Close()

	p1Races := make([]race, 0)

	for i := 1; i < len(times); i++ {
		t, _ := strconv.Atoi(times[i])
		d, _ := strconv.Atoi(distances[i])
		p1Races = append(p1Races, race{t, d})
	}

	t, _ := strconv.Atoi(strings.Join(times[1:], ""))
	d, _ := strconv.Atoi(strings.Join(distances[1:], ""))
	p2Race := race{t, d}

	return &d6{p1Races, p2Race}
}

func (d *d6) part1() int {
	result := 1

	for _, r := range d.p1Races {
		waysToWin := 0

		for bt := 1; bt < r.time; bt++ {
			if findDistance(&r, bt) > r.dist {
				waysToWin++
			}
		}

		result *= waysToWin
	}

	return result
}

func (d *d6) part2() int {
	waysToWin := 0

	for bt := 1; bt < d.p2Race.time; bt++ {
		if findDistance(&d.p2Race, bt) > d.p2Race.dist {
			waysToWin++
		}
	}

	return waysToWin
}

func findDistance(r *race, buttonTime int) int {
	return buttonTime * (r.time - buttonTime)
}
