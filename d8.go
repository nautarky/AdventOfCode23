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

	for steps := 0; ; steps++ {
		if cur == "ZZZ" {
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
