package main

import (
	"bufio"
	"os"
)

type d10 struct {
	pipes []string
}

func newD10(path string) *d10 {
	file, err := os.Open(path)
	check(err)
	defer file.Close()
	s := bufio.NewScanner(file)
	pipes := make([]string, 0)

	for s.Scan() {
		pipes = append(pipes, s.Text())
	}

	return &d10{pipes}
}

func (d *d10) part1() int {
	return 0
}
