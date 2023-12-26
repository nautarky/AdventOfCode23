package main

import (
	"Advent23/day1"
	"Advent23/shared"
	"fmt"
)

func main() {
	input := shared.ReadFile("input/1.txt")
	fmt.Printf("D1P1: %d\n", day1.Part1(input))
	fmt.Printf("D1P2: %d\n", day1.Part2(input))
}
