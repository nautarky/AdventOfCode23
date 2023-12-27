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
