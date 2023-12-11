package main

import (
	"bufio"
	"os"
	"strings"
)

type d8 struct {
	instructions string
	network      map[string][2]string
}

func newD8(path string) *d8 {
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	s := bufio.NewScanner(file)
	s.Scan()
	instructions := s.Text()
	s.Scan()

	network := make(map[string][2]string)

	for s.Scan() {
		parts := strings.Fields(s.Text())
		left := parts[2][1:4]
		right := parts[3][:3]
		network[parts[0]] = [2]string{left, right}
	}

	return &d8{instructions, network}
}

func (d *d8) part1() int {
	cur := "AAA"
	stopCond := func(s string) bool {
		return s == "ZZZ"
	}
	return d.findCycleLen(cur, stopCond)
}

func (d *d8) part2() int {
	curNodes := d.findStartNodes()
	cycles := make([]int, 0)
	stopCond := func(s string) bool {
		return s[len(s)-1] == 'Z'
	}

	for _, n := range curNodes {
		cycles = append(cycles, d.findCycleLen(n, stopCond))
	}

	return LCM(cycles[0], cycles[1], cycles[2:]...)
}

func (d *d8) findCycleLen(node string, stopCond func(string) bool) int {
	cur := node

	for steps := 0; ; steps++ {
		if stopCond(cur) {
			return steps
		}

		direction := d.instructions[steps%len(d.instructions)]

		if direction == 'L' {
			cur = d.network[cur][0]
		} else {
			cur = d.network[cur][1]
		}
	}
}

func (d *d8) findStartNodes() []string {
	nodes := make([]string, 0)

	for k := range d.network {
		if k[len(k)-1] == 'A' {
			nodes = append(nodes, k)
		}
	}

	return nodes
}
