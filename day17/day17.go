package day17

import (
	"container/heap"
	"fmt"
	"math"
	"slices"
)

func Part1(lines []string) int {
	sol := newSolution(lines)
	return sol.dijkstra(complex(0, 0))
}

func Part2(lines []string) int {
	return 0
}

type solution struct {
	tiles   map[complex64]int
	visited map[complex64]bool
	items   map[complex64]*itemComplex64
	costs   priorityQueueItemComplex64
	dest    complex64
}

func (s *solution) dijkstra(src complex64) int {
	c := heap.Pop(&s.costs).(*itemComplex64)
	neighbors := []complex64{complex(1, 0), complex(0, 1)}
	prev := complex64(complex(-1, -1))

	for {
		for _, n := range neighbors {
			tile := s.tiles[n]
			item := s.items[n]

			if c.priority+tile < item.priority {
				s.costs.Update(item, c.priority+tile)
			}

			fmt.Println(c)
			fmt.Println(item)
			fmt.Println("----------")
		}

		if c.value == s.dest {
			return c.priority
		}
		s.visited[c.value] = true

		prev = c.value
		c = heap.Pop(&s.costs).(*itemComplex64)
		neighbors = s.unvisitedNeighbors(c.value, prev)
	}
}

func (s *solution) unvisitedNeighbors(node, prev complex64) []complex64 {
	neighbors := make([]complex64, 0)
	diff := node - prev

	if real(diff) != 0 {
		// going vertically
		neighbors = append(neighbors, node+complex(1, 0), node+complex(-1, 0))
	} else {
		// going horizontally
		neighbors = append(neighbors, node+complex(0, 1), node+complex(0, -1))
	}

	return slices.DeleteFunc(neighbors, func(n complex64) bool {
		visited, ok := s.visited[n]
		return visited || !ok
	})
}

func newSolution(lines []string) solution {
	tiles := make(map[complex64]int)
	items := make(map[complex64]*itemComplex64)
	costs := make(priorityQueueItemComplex64, 0)
	visited := make(map[complex64]bool)

	for i, line := range lines {
		for j, val := range line {
			value := complex(float32(i), float32(j))
			tiles[value] = int(val - '0')

			var item *itemComplex64
			if i == 0 && j == 0 {
				visited[complex(float32(i), float32(j))] = true
				item = &itemComplex64{value: value, priority: 0}
			} else {
				visited[complex(float32(i), float32(j))] = false
				item = &itemComplex64{value: value, priority: math.MaxInt}
			}

			items[value] = item
			heap.Push(&costs, item)
		}
	}

	return solution{tiles, visited, items, costs, complex(float32(len(lines)-1), float32(len(lines[0])-1))}
}

// An itemComplex64 is something we manage in a priority queue.
type itemComplex64 struct {
	value     complex64 // The value of the item; arbitrary.
	direction complex64
	streak    int
	priority  int // The priority of the item in the queue.
	index     int // The index of the item in the heap. Needed by update and is maintained by the heap.Interface methods.
}

// A priorityQueueItemComplex64 implements heap.Interface and holds Items.
type priorityQueueItemComplex64 []*itemComplex64

func (pq priorityQueueItemComplex64) Len() int { return len(pq) }

func (pq priorityQueueItemComplex64) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority < pq[j].priority
}

func (pq priorityQueueItemComplex64) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueueItemComplex64) Push(x any) {
	n := len(*pq)
	item := x.(*itemComplex64)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *priorityQueueItemComplex64) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority of an Item in the queue.
func (pq *priorityQueueItemComplex64) Update(item *itemComplex64, priority int) {
	item.priority = priority
	heap.Fix(pq, item.index)
}
