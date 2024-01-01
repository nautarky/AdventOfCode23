package day15

import (
	"slices"
	"strconv"
	"strings"
)

func Part1(lines []string) int {
	sum := 0
	for _, s := range strings.Split(lines[0], ",") {
		sum += hash(s)
	}
	return sum
}

func Part2(lines []string) int {
	boxes := make([]box, 256)
	parts := strings.Split(lines[0], ",")

	for _, p := range parts {
		opPos := strings.IndexAny(p, "=-")
		op := p[opPos]

		label := p[:opPos]
		b := hash(label)

		if op == '=' {
			fl, _ := strconv.Atoi(p[opPos+1:])
			s := lens{label, fl}

			i := slices.IndexFunc(boxes[b], func(l lens) bool { return l.label == label })
			if i == -1 {
				boxes[b] = append(boxes[b], s)
			} else {
				boxes[b][i] = s
			}
		} else {
			boxes[b] = slices.DeleteFunc(boxes[b], func(l lens) bool { return l.label == label })
		}
	}

	sum := 0

	for bi, b := range boxes {
		for li, l := range b {
			sum += (bi + 1) * (li + 1) * l.focal
		}
	}

	return sum
}

type box []lens

type lens struct {
	label string
	focal int
}

func hash(s string) int {
	cur := 0
	for _, b := range s {
		cur = ((cur + int(b)) * 17) % 256
	}
	return cur
}
