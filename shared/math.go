package shared

import "math"

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
