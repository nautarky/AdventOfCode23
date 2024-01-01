package day10

import (
	"Advent23/shared"
	"errors"
)

func Part1(lines []string) int {
	g := newGrid(lines)
	start, _ := g.findStart()
	paths, openings := g.findAdjacentPaths(start)
	g.replaceStart(start, openings)

	cur := paths[0]
	prev := start
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

func Part2(lines []string) int {
	g := newGrid(lines)
	start, _ := g.findStart()
	paths, openings := g.findAdjacentPaths(start)
	g.replaceStart(start, openings)

	points := []shared.Point{start}
	cur := paths[0]
	prev := start

	for cur != start {
		paths = g.findPathsOut(cur)

		if paths[0] == prev {
			cur, prev = paths[1], cur
		} else {
			cur, prev = paths[0], cur
		}

		points = append(points, cur)
	}

	area := shared.ShoelaceTheorem(points)
	picks := shared.PicksTheoremI(area, float64(len(points)))
	return int(picks)
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

func (g *grid) findStart() (shared.Point, error) {
	for i, row := range g.tiles {
		for j, tile := range row {
			if tile == 'S' {
				return shared.Point{X: j, Y: i}, nil
			}
		}
	}
	return shared.Point{}, errors.New("didn't find a start position")
}

func (g *grid) findPathsOut(p shared.Point) []shared.Point {
	switch g.tiles[p.Y][p.X] {
	case '|':
		return []shared.Point{{X: p.X, Y: p.Y - 1}, {X: p.X, Y: p.Y + 1}}
	case '-':
		return []shared.Point{{X: p.X - 1, Y: p.Y}, {X: p.X + 1, Y: p.Y}}
	case 'L':
		return []shared.Point{{X: p.X, Y: p.Y - 1}, {X: p.X + 1, Y: p.Y}}
	case 'J':
		return []shared.Point{{X: p.X, Y: p.Y - 1}, {X: p.X - 1, Y: p.Y}}
	case '7':
		return []shared.Point{{X: p.X - 1, Y: p.Y}, {X: p.X, Y: p.Y + 1}}
	case 'F':
		return []shared.Point{{X: p.X + 1, Y: p.Y}, {X: p.X, Y: p.Y + 1}}
	}

	return []shared.Point{}
}

func (g *grid) findAdjacentPaths(p shared.Point) ([]shared.Point, map[rune]bool) {
	adjacent := make([]shared.Point, 0)
	directions := map[rune]bool{
		'n': false,
		'e': false,
		's': false,
		'w': false,
	}

	// north
	if p.Y-1 > -1 {
		t := g.tiles[p.Y-1][p.X]

		if t == '|' || t == '7' || t == 'F' {
			adjacent = append(adjacent, shared.Point{X: p.X, Y: p.Y - 1})
			directions['n'] = true
		}
	}

	// east
	if p.X+1 < len(g.tiles[p.Y]) {
		t := g.tiles[p.Y][p.X+1]

		if t == '-' || t == 'J' || t == '7' {
			adjacent = append(adjacent, shared.Point{X: p.X + 1, Y: p.Y})
			directions['e'] = true
		}
	}

	// south
	if p.Y+1 < len(g.tiles) {
		t := g.tiles[p.Y+1][p.X]

		if t == '|' || t == 'J' || t == 'L' {
			adjacent = append(adjacent, shared.Point{X: p.X, Y: p.Y + 1})
			directions['s'] = true
		}
	}

	// west
	if p.X-1 > -1 {
		t := g.tiles[p.Y][p.X-1]

		if t == '-' || t == 'L' || t == 'F' {
			adjacent = append(adjacent, shared.Point{X: p.X - 1, Y: p.Y})
			directions['w'] = true
		}
	}

	return adjacent, directions
}

func (g *grid) replaceStart(start shared.Point, openings map[rune]bool) {
	switch {
	case openings['n'] && openings['e']:
		g.tiles[start.Y][start.X] = 'L'
	case openings['n'] && openings['s']:
		g.tiles[start.Y][start.X] = '|'
	case openings['n'] && openings['w']:
		g.tiles[start.Y][start.X] = 'J'
	case openings['s'] && openings['e']:
		g.tiles[start.Y][start.X] = 'F'
	case openings['s'] && openings['w']:
		g.tiles[start.Y][start.X] = '7'
	case openings['w'] && openings['e']:
		g.tiles[start.Y][start.X] = '-'
	}
}
