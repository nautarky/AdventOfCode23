package main

import (
	"bufio"
	"os"
	"strings"
	"unicode"
)

const d1InputPath = "input/1.txt"

var atoiMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

// Returns value, index
func getLeftPart1(line string) (int, int) {
	for pos, char := range line {
		if unicode.IsDigit(char) {
			return int(char - '0'), pos
		}
	}

	return -1, -1
}

func getLeftPart2(line string) int {
	digit, lowIndex := getLeftPart1(line)

	for k := range atoiMap {
		i := strings.Index(line, k)
		if i != -1 && i < lowIndex {
			digit = atoiMap[k]
			lowIndex = i
		}
	}

	return digit
}

// Returns value, index
func getRightPart1(line string) (int, int) {
	runes := []rune(line)

	for i := len(runes) - 1; i > -1; i-- {
		if unicode.IsDigit(runes[i]) {
			return int(runes[i] - '0'), i
		}
	}

	return -1, -1
}

func getRightPart2(line string) int {
	digit, highIndex := getRightPart1(line)

	for k := range atoiMap {
		i := strings.LastIndex(line, k)
		if i != -1 && i > highIndex {
			digit = atoiMap[k]
			highIndex = i
		}
	}

	return digit
}

func d1() (int, int) {
	file, err := os.Open(d1InputPath)
	check(err)
	defer file.Close()

	part1Sum, part2Sum := 0, 0
	s := bufio.NewScanner(file)
	for s.Scan() {
		line := s.Text()
		l1, _ := getLeftPart1(line)
		r1, _ := getRightPart1(line)
		l2 := getLeftPart2(line)
		r2 := getRightPart2(line)
		part1Sum += l1*10 + r1
		part2Sum += l2*10 + r2
	}

	return part1Sum, part2Sum
}
