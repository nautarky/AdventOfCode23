package day16

import (
	"Advent23/shared"
	"slices"
)

func Part1(lines []string) int {
	lights := []light{{p: shared.Point{X: 0, Y: 0}, v: shared.UnitVector{X: 1, Y: 0}}}
	visited := make(map[light]bool)
	energized := make(map[shared.Point]bool)

	for len(lights) > 0 {
		curLight := lights[0]
		lights = lights[1:]

		seen := visited[curLight]
		if seen {
			continue
		}

		visited[curLight] = true
		energized[curLight.p] = true

		lights = append(lights, getNewLights(lines, curLight)...)
	}

	return len(energized)
}

func Part2(lines []string) int {
	return 0
}

type light struct {
	p shared.Point
	v shared.UnitVector
}

func getNewLights(lines []string, l light) []light {
	lights := make([]light, 0)
	tile := lines[l.p.Y][l.p.X]

	if tile == '.' {
		newPos := shared.ApplyUnitVector(l.p, l.v)
		lights = append(lights, light{newPos, l.v})
	}

	if tile == '\\' || tile == '/' {
		lights = append(lights, reflectLight(l, tile))
	}

	if tile == '|' || tile == '-' {
		lights = append(lights, splitLight(l, tile)...)
	}

	return slices.DeleteFunc(lights, func(l light) bool { return !isValidPoint(lines, l.p) })
}

// reflectLight does not have tile context, so it may return invalid lights
func reflectLight(l light, mirror byte) light {
	var v shared.UnitVector

	if mirror == '\\' {
		switch l.v {
		case shared.UnitVector{X: 1, Y: 0}:
			v = shared.UnitVector{X: 0, Y: 1}
		case shared.UnitVector{X: -1, Y: 0}:
			v = shared.UnitVector{X: 0, Y: -1}
		case shared.UnitVector{X: 0, Y: 1}:
			v = shared.UnitVector{X: 1, Y: 0}
		default:
			v = shared.UnitVector{X: -1, Y: 0}
		}
	} else {
		switch l.v {
		case shared.UnitVector{X: 1, Y: 0}:
			v = shared.UnitVector{X: 0, Y: -1}
		case shared.UnitVector{X: -1, Y: 0}:
			v = shared.UnitVector{X: 0, Y: 1}
		case shared.UnitVector{X: 0, Y: 1}:
			v = shared.UnitVector{X: -1, Y: 0}
		default:
			v = shared.UnitVector{X: 1, Y: 0}
		}
	}

	return light{shared.ApplyUnitVector(l.p, v), v}
}

// splitLight does not have tile context, so it may return invalid lights
func splitLight(l light, splitter byte) []light {
	vecs := make([]shared.UnitVector, 0)

	if splitter == '-' {
		if l.v.X == -1 || l.v.X == 1 {
			vecs = append(vecs, l.v)
		} else {
			v1 := shared.UnitVector{X: 1, Y: 0}
			v2 := shared.UnitVector{X: -1, Y: 0}
			vecs = append(vecs, v1, v2)
		}
	} else {
		if l.v.Y == -1 || l.v.Y == 1 {
			vecs = append(vecs, l.v)
		} else {
			v1 := shared.UnitVector{X: 0, Y: 1}
			v2 := shared.UnitVector{X: 0, Y: -1}
			vecs = append(vecs, v1, v2)
		}
	}

	lights := make([]light, len(vecs))
	for i, vec := range vecs {
		lights[i] = light{shared.ApplyUnitVector(l.p, vec), vec}
	}

	return lights
}

func isValidPoint(lines []string, p shared.Point) bool {
	return p.Y > -1 && p.Y < len(lines) && p.X > -1 && p.X < len(lines[0])
}
