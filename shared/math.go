package shared

import (
	"container/heap"
	"math"
)

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func QuadraticFormula(a, b, c float64) (float64, float64) {
	plus := (-b + math.Sqrt(b*b-4*a*c)) / (2 * a)
	minus := (-b - math.Sqrt(b*b-4*a*c)) / (2 * a)
	return plus, minus
}

type Point struct {
	X int
	Y int
}

type UnitVector struct {
	X int
	Y int
}

// ApplyUnitVector assumes 2d matrix positions. N->S: N+1
func ApplyUnitVector(p Point, v UnitVector) Point {
	return Point{p.X + v.X, p.Y + v.Y}
}

func ShoelaceTheorem(points []Point) float64 {
	sum := 0.0
	for i := 0; i < len(points); i++ {
		p1, p2 := points[i], points[(i+1)%len(points)]
		sum += float64(p1.X*p2.Y) - float64(p1.Y*p2.X)
	}

	return math.Abs(sum / 2.0)
}

// PicksTheoremI solves for I
func PicksTheoremI(a, b float64) float64 {
	return -1 * (-a + (b / 2) - 1)
}

func SumIntSlice(ints []int) int {
	sum := 0

	for _, i := range ints {
		sum += i
	}

	return sum
}

// An ItemComplex64 is something we manage in a priority queue.
type ItemComplex64 struct {
	Value    complex64 // The value of the item; arbitrary.
	Priority int       // The priority of the item in the queue.
	// The Index is needed by update and is maintained by the heap.Interface methods.
	Index int // The index of the item in the heap.
}

// A PriorityQueueItemComplex64 implements heap.Interface and holds Items.
type PriorityQueueItemComplex64 []*ItemComplex64

func (pq PriorityQueueItemComplex64) Len() int { return len(pq) }

func (pq PriorityQueueItemComplex64) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].Priority > pq[j].Priority
}

func (pq PriorityQueueItemComplex64) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueueItemComplex64) Push(x any) {
	n := len(*pq)
	item := x.(*ItemComplex64)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueueItemComplex64) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.Index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueueItemComplex64) update(item *ItemComplex64, value complex64, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, item.Index)
}
