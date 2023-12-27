package day7

import (
	"sort"
	"strconv"
	"strings"
)

func Part1(lines []string) int {
	return solve(lines, false)
}

func Part2(lines []string) int {
	return solve(lines, true)
}

func solve(lines []string, isP2 bool) int {
	hands := make([]hand, len(lines))

	for i, line := range lines {
		hands[i] = buildHand(line, isP2)
	}

	rankHands(hands)

	sum := 0
	for i, h := range hands {
		sum += (i + 1) * h.bid
	}

	return sum
}

func rankHands(hands []hand) {
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].typ == hands[j].typ {
			for v := range hands[i].values {
				if hands[i].values[v] != hands[j].values[v] {
					return hands[i].values[v] < hands[j].values[v]
				}
			}
		}

		return hands[i].typ < hands[j].typ
	})
}

var p1CardValues = map[byte]int{'2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9, 'T': 10, 'J': 11, 'Q': 12, 'K': 13, 'A': 14}
var p2CardValues = map[byte]int{'2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9, 'T': 10, 'J': 1, 'Q': 12, 'K': 13, 'A': 14}

type handType int

const (
	highCard handType = iota
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

type hand struct {
	typ    handType
	bid    int
	values []int
}

func buildHand(line string, isP2 bool) hand {
	counts := make(map[rune]int)
	parts := strings.Fields(line)
	bid, _ := strconv.Atoi(parts[1])

	for _, r := range parts[0] {
		counts[r] += 1
	}

	var most, next int
	for k, v := range counts {
		if isP2 && k == 'J' {
			continue
		}

		if v >= most {
			most, next = v, most
		} else if v > next {
			next = v
		}
	}

	if isP2 {
		most += counts['J']
	}

	typ := findHandType(most, next)

	values := make([]int, 5)
	for i := range values {
		if isP2 {
			values[i] = p2CardValues[parts[0][i]]
		} else {
			values[i] = p1CardValues[parts[0][i]]
		}
	}

	return hand{typ, bid, values}
}

func findHandType(most, next int) handType {
	var typ handType

	switch most {
	case 5:
		typ = fiveOfAKind
	case 4:
		typ = fourOfAKind
	case 3:
		if next == 2 {
			typ = fullHouse
		} else {
			typ = threeOfAKind
		}
	case 2:
		if next == 2 {
			typ = twoPair
		} else {
			typ = onePair
		}
	default:
		typ = highCard
	}

	return typ
}
