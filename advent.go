package main

import (
	"Advent23/day1"
	"Advent23/shared"
	"fmt"
)

func main() {
	input := shared.ReadFile("input/1_test.txt")
	fmt.Printf("D1P1: %d\n", day1.Part1(input))
}
