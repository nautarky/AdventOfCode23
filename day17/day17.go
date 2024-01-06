package day17

import (
	"Advent23/shared"
	"container/heap"
)

func Part1(lines []string) int {
	sol := newSolution(lines)
	return 0
}

func Part2(lines []string) int {
	return 0
}

type solution struct {
	tiles   map[complex64]int
	visited map[complex64]bool
	items   map[complex64]*shared.ItemComplex64
	costs   shared.PriorityQueueItemComplex64
}

func newSolution(lines []string) solution {
	tiles := make(map[complex64]int)
	items := make(map[complex64]*shared.ItemComplex64)
	costs := make(shared.PriorityQueueItemComplex64, 0)

	for i, line := range lines {
		for j, val := range line {
			cost := int(val - '0')
			value := complex(float32(i), float32(j))

			var item *shared.ItemComplex64
			if i == 0 && j == 0 {
				item = &shared.ItemComplex64{Value: value, Priority: 0}
			} else {
				item = &shared.ItemComplex64{Value: value, Priority: cost}
			}

			items[value] = item
			heap.Push(&costs, item)
		}
	}

	visited := make(map[complex64]bool)
	return solution{tiles, visited, items, costs}
}
