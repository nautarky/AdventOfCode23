package main

import (
	"Advent23/day01"
	"Advent23/day02"
	"Advent23/day03"
	"Advent23/day04"
	"Advent23/day05"
	"Advent23/day06"
	"Advent23/day07"
	"Advent23/day08"
	"Advent23/day09"
	"Advent23/day10"
	"Advent23/day11"
	"Advent23/day13"
	"Advent23/day14"
	"Advent23/day15"
	"Advent23/day18"
	"Advent23/shared"
	"fmt"
)

func main() {
	printOutput("input/1.txt", day01.Part1, day01.Part2, 1)
	printOutput("input/2.txt", day02.Part1, day02.Part2, 2)
	printOutput("input/3.txt", day03.Part1, day03.Part2, 3)
	printOutput("input/4.txt", day04.Part1, day04.Part2, 4)
	printOutput("input/5.txt", day05.Part1, day05.Part2, 5)
	printOutput("input/6.txt", day06.Part1, day06.Part2, 6)
	printOutput("input/7.txt", day07.Part1, day07.Part2, 7)
	printOutput("input/8.txt", day08.Part1, day08.Part2, 8)
	printOutput("input/9.txt", day09.Part1, day09.Part2, 9)
	printOutput("input/10.txt", day10.Part1, day10.Part2, 10)
	printOutput("input/11.txt", day11.Part1, day11.Part2, 11)
	// printOutput("input/12.txt", day12.Part1, day12.Part2, 12)
	printOutput("input/13.txt", day13.Part1, day13.Part2, 13)
	printOutput("input/14.txt", day14.Part1, day14.Part2, 14)
	printOutput("input/15.txt", day15.Part1, day15.Part2, 15)
	// printOutput("input/16.txt", day16.Part1, day16.Part2, 16)
	// printOutput("input/17.txt", day17.Part1, day17.Part2, 17)
	printOutput("input/18.txt", day18.Part1, day18.Part2, 18)
}

type solution func([]string) int

func printOutput(path string, part1, part2 solution, day int) {
	input := shared.ReadFile(path)
	fmt.Printf("Day %02d\n", day)
	fmt.Printf("  Part 1: %d\n", part1(input))
	fmt.Printf("  Part 2: %d\n", part2(input))
}
