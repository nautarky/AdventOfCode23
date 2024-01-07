package day17

import (
	"container/heap"
	"math"
)

func Part1(lines []string) int {
	sol := newSolution(lines, 1, 3)
	return sol.dijkstra()
}

func Part2(lines []string) int {
	sol := newSolution(lines, 4, 10)
	return sol.dijkstra()
}

type solution struct {
	tiles      map[complex64]int
	visited    map[itemKey]bool
	items      map[itemKey]*queueItem
	costs      priorityQueueItemComplex64
	dest       complex64
	directions []complex64
	minStreak  int
	maxStreak  int
}

func (s *solution) dijkstra() int {
	for {
		c := heap.Pop(&s.costs).(*queueItem)
		neighbors := s.unvisitedNeighbors(c.itemKey)

		for _, n := range neighbors {
			tile := s.tiles[n.value]
			item := s.items[n]

			if c.priority+tile < item.priority {
				s.costs.Update(item, c.priority+tile)
			}
		}

		if c.value == s.dest && c.streak >= s.minStreak {
			return c.priority
		}
		s.visited[c.itemKey] = true
	}
}

func (s *solution) unvisitedNeighbors(node itemKey) []itemKey {
	neighbors := make([]itemKey, 0)

	for _, d := range s.directions {
		// find position of moving in this direction
		value := node.value + d

		// dont go out of bounds or backwards
		_, ok := s.tiles[value]
		if !ok || d == -1*node.direction {
			continue
		}

		if d != node.direction && node.streak < s.minStreak {
			continue
		}

		streak := 1
		if d == node.direction {
			streak = node.streak + 1
		}

		key := itemKey{value, d, streak}
		seen := s.visited[key]

		if streak > s.maxStreak || seen {
			continue
		}

		neighbors = append(neighbors, itemKey{node.value + d, d, streak})
	}

	return neighbors
}

func newSolution(lines []string, minStreak, maxStreak int) solution {
	tiles := make(map[complex64]int)
	items := make(map[itemKey]*queueItem)
	costs := make(priorityQueueItemComplex64, 0)
	visited := make(map[itemKey]bool)
	directions := []complex64{complex(-1, 0), complex(1, 0), complex(0, -1), complex(0, 1)}

	for i, line := range lines {
		for j, val := range line {
			value := complex(float32(i), float32(j))
			tiles[value] = int(val - '0')

			for _, d := range directions {
				for streak := 1; streak < maxStreak+1; streak++ {
					key := itemKey{value, d, streak}
					item := &queueItem{itemKey: key, priority: math.MaxInt}
					visited[key] = false
					if key.value == 0 && streak == 1 {
						item.priority = 0
						visited[key] = true
					}

					items[key] = item
					heap.Push(&costs, item)
				}
			}
		}
	}

	return solution{tiles, visited, items, costs, complex(float32(len(lines)-1), float32(len(lines[0])-1)), directions, minStreak, maxStreak}
}

type itemKey struct {
	value     complex64
	direction complex64
	streak    int
}

// An queueItem is something we manage in a priority queue.
type queueItem struct {
	itemKey
	priority int // The priority of the item in the queue.
	index    int // The index of the item in the heap. Needed by update and is maintained by the heap.Interface methods.
}

// A priorityQueueItemComplex64 implements heap.Interface and holds Items.
type priorityQueueItemComplex64 []*queueItem

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
	item := x.(*queueItem)
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
func (pq *priorityQueueItemComplex64) Update(item *queueItem, priority int) {
	item.priority = priority
	heap.Fix(pq, item.index)
}
