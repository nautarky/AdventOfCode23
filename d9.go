package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type d9 struct {
	reports [][]int
}

func newD9(path string) *d9 {
	file, err := os.Open(path)
	check(err)
	defer file.Close()
	s := bufio.NewScanner(file)

	reports := make([][]int, 0)
	for s.Scan() {
		parts := strings.Fields(s.Text())
		report := make([]int, len(parts))

		for i, p := range parts {
			num, _ := strconv.Atoi(p)
			report[i] = num
		}

		reports = append(reports, report)
	}

	return &d9{reports}
}

func (d *d9) part1() int {
	sum := 0

	for _, r := range d.reports {
		sum += predictNext(r)
	}

	return sum
}

func (d *d9) part2() int {
	sum := 0

	for _, r := range d.reports {
		sum += predictPrev(r)
	}

	return sum
}

func predictNext(report []int) int {
	if isBaseCase(report) {
		return 0
	}

	nextLevel := makeNewLevel(report)
	return report[len(report)-1] + predictNext(nextLevel)
}

func predictPrev(report []int) int {
	if isBaseCase(report) {
		return 0
	}

	nextLevel := makeNewLevel(report)
	return report[0] - predictPrev(nextLevel)
}

func isBaseCase(report []int) bool {
	for _, r := range report {
		if r != 0 {
			return false
		}
	}

	return true
}

func makeNewLevel(report []int) []int {
	newLevel := make([]int, len(report)-1)

	for i := 0; i < len(report)-1; i++ {
		newLevel[i] = report[i+1] - report[i]
	}

	return newLevel
}
