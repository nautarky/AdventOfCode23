package main

import (
	"bufio"
	"os"
	"strings"
)

type d4 struct {
	input []string
}

func newD4(path string) *d4 {
	file, err := os.Open(path)
	check(err)

	input := make([]string, 0)
	s := bufio.NewScanner(file)
	for s.Scan() {
		input = append(input, s.Text())
	}

	err = file.Close()
	check(err)

	return &d4{input}
}

func (d *d4) part1() int {
	sum := 0

	for _, s := range d.input {
		winningSet, mySet := buildCardSets(s)

		matches := findMatches(winningSet, mySet)
		if matches == 0 {
			continue
		}

		score := 1
		for matches > 1 {
			score *= 2
			matches--
		}

		sum += score
	}

	return sum
}

func (d *d4) part2() int {
	counts := make([]int, len(d.input))

	for i := range counts {
		winningSet, mySet := buildCardSets(d.input[i])
		matches := findMatches(winningSet, mySet)

		for j := 1; j < matches+1; j++ {
			counts[i+j] += counts[i] + 1
		}
	}

	sum := len(counts)
	for _, count := range counts {
		sum += count
	}

	return sum
}

func buildCardSets(s string) (map[string]struct{}, map[string]struct{}) {
	parts := strings.Fields(s)
	delim := 0

	for i := 2; i < len(parts); i++ {
		if parts[i] == "|" {
			delim = i
			break
		}
	}

	return buildSet(parts[2:delim]), buildSet(parts[delim+1:])
}

func findMatches(a, b map[string]struct{}) int {
	matches := 0

	for k, _ := range a {
		_, ok := b[k]
		if ok {
			matches++
		}
	}

	return matches
}

func buildSet(elems []string) map[string]struct{} {
	set := make(map[string]struct{})

	for _, s := range elems {
		set[s] = struct{}{}
	}

	return set
}
