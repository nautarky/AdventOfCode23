package main

import (
	"Advent23/day1"
	"Advent23/day2"
	"Advent23/day3"
	"Advent23/day4"
	"Advent23/shared"
	"fmt"
)

func main() {
	printOutput("input/1.txt", day1.Part1, day1.Part2, 1)
	printOutput("input/2.txt", day2.Part1, day2.Part2, 2)
	printOutput("input/3.txt", day3.Part1, day3.Part2, 3)
	printOutput("input/4.txt", day4.Part1, day4.Part2, 4)
}

type solution func([]string) int

func printOutput(path string, part1, part2 solution, day int) {
	input := shared.ReadFile(path)
	fmt.Printf("Day %02d\n", day)
	fmt.Printf("  Part 1: %d\n", part1(input))
	fmt.Printf("  Part 2: %d\n", part2(input))
}
