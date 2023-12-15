package main

import (
	"bufio"
	"errors"
	"os"
)

type point struct {
	x int
	y int
}

func (p *point) isEqual(compare point) bool {
	return p.x == compare.x && p.y == compare.y
}

type d10 struct {
	pipes []string
}

func newD10(path string) *d10 {
	file, err := os.Open(path)
	check(err)
	defer file.Close()
	s := bufio.NewScanner(file)
	pipes := make([]string, 0)

	for s.Scan() {
		pipes = append(pipes, s.Text())
	}

	return &d10{pipes}
}

func (d *d10) part1() int {
	//start := d.findStart()
	//exits, _ := d.findExits(start)
	return 0
}

func (d *d10) findStart() point {
	for i, row := range d.pipes {
		for j, pipe := range row {
			if pipe == 'S' {
				return point{j, i}
			}
		}
	}

	return point{-1, -1}
}

func (d *d10) findExits(p point) ([2]point, error) {
	if p.y < 0 || p.y > len(d.pipes) || p.x < 0 || p.x > len(d.pipes[p.y]) {
		return [2]point{}, errors.New("not a valid point")
	}

	pipe := d.pipes[p.y][p.x]

	north := point{p.x, p.y - 1}
	south := point{p.x, p.y + 1}
	east := point{p.x + 1, p.y}
	west := point{p.x - 1, p.y}

	switch pipe {
	case 'S':
		// Handle starting position as a special case
		return d.findStartExits(p, []point{north, south, east, west}), nil
	case '|':
		return [2]point{north, south}, nil
	case '-':
		return [2]point{east, west}, nil
	case 'L':
		return [2]point{north, east}, nil
	case 'J':
		return [2]point{north, west}, nil
	case '7':
		return [2]point{south, west}, nil
	case 'F':
		return [2]point{south, east}, nil
	default:
		return [2]point{{-1, -1}, {-1, -1}}, nil
	}
}

func (d *d10) findStartExits(p point, candidates []point) [2]point {
	i, result := 0, [2]point{}

	for _, c := range candidates {
		exits, err := d.findExits(c)
		if err != nil {
			continue
		}

		for _, e := range exits {
			if e.isEqual(p) {
				result[i] = c
				i++
				break
			}
		}
	}

	return result
}
