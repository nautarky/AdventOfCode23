package day10

import (
	"errors"
	"fmt"
)

func Part1(lines []string) int {
	g := newGrid(lines)
	start, _ := g.findStart()
	paths, openings := g.findAdjacentPaths(start)
	g.replaceStart(start, openings)

	cur, prev := paths[0], start
	steps := 1

	for cur != start {
		paths = g.findPathsOut(cur)

		if paths[0] == prev {
			cur, prev = paths[1], cur
		} else {
			cur, prev = paths[0], cur
		}

		steps++
	}

	return steps / 2
}

func debug(tiles [][]rune, cur point) {
	t := tiles[cur.y][cur.x]

	tiles[cur.y][cur.x] = '*'

	fmt.Println("--------------------------------")
	for _, row := range tiles {
		fmt.Println(string(row))
	}
	fmt.Println("--------------------------------")

	tiles[cur.y][cur.x] = t
}

func Part2(lines []string) int {
	return 0
}

type point struct {
	x int
	y int
}

type grid struct {
	tiles [][]rune
}

func newGrid(lines []string) *grid {
	tiles := make([][]rune, len(lines))

	for i, line := range lines {
		tiles[i] = []rune(line)
	}

	return &grid{tiles}
}

func (g *grid) findStart() (point, error) {
	for i, row := range g.tiles {
		for j, tile := range row {
			if tile == 'S' {
				return point{j, i}, nil
			}
		}
	}
	return point{}, errors.New("didn't find a start position")
}

func (g *grid) findPathsOut(p point) []point {
	switch g.tiles[p.y][p.x] {
	case '|':
		return []point{{p.x, p.y - 1}, {p.x, p.y + 1}}
	case '-':
		return []point{{p.x - 1, p.y}, {p.x + 1, p.y}}
	case 'L':
		return []point{{p.x, p.y - 1}, {p.x + 1, p.y}}
	case 'J':
		return []point{{p.x, p.y - 1}, {p.x - 1, p.y}}
	case '7':
		return []point{{p.x - 1, p.y}, {p.x, p.y + 1}}
	case 'F':
		return []point{{p.x + 1, p.y}, {p.x, p.y + 1}}
	}

	return []point{}
}

func (g *grid) findAdjacentPaths(p point) ([]point, map[rune]bool) {
	adjacent := make([]point, 0)
	directions := map[rune]bool{
		'n': false,
		'e': false,
		's': false,
		'w': false,
	}

	// north
	if p.y-1 > -1 {
		t := g.tiles[p.y-1][p.x]

		if t == '|' || t == '7' || t == 'F' {
			adjacent = append(adjacent, point{p.x, p.y - 1})
			directions['n'] = true
		}
	}

	// east
	if p.x+1 < len(g.tiles[p.y]) {
		t := g.tiles[p.y][p.x+1]

		if t == '-' || t == 'J' || t == '7' {
			adjacent = append(adjacent, point{p.x + 1, p.y})
			directions['e'] = true
		}
	}

	// south
	if p.y+1 < len(g.tiles) {
		t := g.tiles[p.y+1][p.x]

		if t == '|' || t == 'J' || t == 'L' {
			adjacent = append(adjacent, point{p.x, p.y + 1})
			directions['s'] = true
		}
	}

	// west
	if p.x-1 > -1 {
		t := g.tiles[p.y][p.x-1]

		if t == '-' || t == 'L' || t == 'F' {
			adjacent = append(adjacent, point{p.x - 1, p.y})
			directions['w'] = true
		}
	}

	return adjacent, directions
}

func (g *grid) replaceStart(start point, openings map[rune]bool) {
	switch {
	case openings['n'] && openings['e']:
		g.tiles[start.y][start.x] = 'L'
	case openings['n'] && openings['s']:
		g.tiles[start.y][start.x] = '|'
	case openings['n'] && openings['w']:
		g.tiles[start.y][start.x] = 'J'
	case openings['s'] && openings['e']:
		g.tiles[start.y][start.x] = 'F'
	case openings['s'] && openings['w']:
		g.tiles[start.y][start.x] = '7'
	case openings['w'] && openings['e']:
		g.tiles[start.y][start.x] = '-'
	}
}
