package main

import (
	"Advent23/day1"
	"Advent23/day10"
	"Advent23/day11"
	"Advent23/day13"
	"Advent23/day14"
	"Advent23/day15"
	"Advent23/day2"
	"Advent23/day3"
	"Advent23/day4"
	"Advent23/day5"
	"Advent23/day6"
	"Advent23/day7"
	"Advent23/day8"
	"Advent23/day9"
	"Advent23/shared"
	"fmt"
)

func main() {
	printOutput("input/1.txt", day1.Part1, day1.Part2, 1)
	printOutput("input/2.txt", day2.Part1, day2.Part2, 2)
	printOutput("input/3.txt", day3.Part1, day3.Part2, 3)
	printOutput("input/4.txt", day4.Part1, day4.Part2, 4)
	printOutput("input/5.txt", day5.Part1, day5.Part2, 5)
	printOutput("input/6.txt", day6.Part1, day6.Part2, 6)
	printOutput("input/7.txt", day7.Part1, day7.Part2, 7)
	printOutput("input/8.txt", day8.Part1, day8.Part2, 8)
	printOutput("input/9.txt", day9.Part1, day9.Part2, 9)
	printOutput("input/10.txt", day10.Part1, day10.Part2, 10)
	printOutput("input/11.txt", day11.Part1, day11.Part2, 11)
	// printOutput("input/12.txt", day12.Part1, day12.Part2, 12)
	printOutput("input/13.txt", day13.Part1, day13.Part2, 13)
	printOutput("input/14.txt", day14.Part1, day14.Part2, 14)
	printOutput("input/15.txt", day15.Part1, day15.Part2, 15)
	// printOutput("input/16.txt", day16.Part1, day16.Part2, 16)
	// printOutput("input/17.txt", day17.Part1, day17.Part2, 17)
}

type solution func([]string) int

func printOutput(path string, part1, part2 solution, day int) {
	input := shared.ReadFile(path)
	fmt.Printf("Day %02d\n", day)
	fmt.Printf("  Part 1: %d\n", part1(input))
	fmt.Printf("  Part 2: %d\n", part2(input))
}
