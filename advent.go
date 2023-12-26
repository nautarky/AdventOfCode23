package main

import (
	"Advent23/day1"
	"Advent23/day2"
	"Advent23/shared"
	"fmt"
)

func main() {
	printOutput("input/1.txt", day1.Part1, day1.Part2, 1)
	printOutput("input/2.txt", day2.Part1, day2.Part2, 2)
}

type solution func([]string) int

func printOutput(path string, part1, part2 solution, day int) {
	input := shared.ReadFile(path)
	fmt.Printf("D%02dP1: %d\n", day, part1(input))
	fmt.Printf("D%02dP2: %d\n", day, part2(input))
}
