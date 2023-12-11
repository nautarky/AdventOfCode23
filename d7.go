package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

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
	cards []int
	typ   handType
	bid   int
}

type d7 struct {
	cardScores map[rune]int
	hands      []hand
}

func newD7(path string) *d7 {
	file, err := os.Open(path)
	check(err)

	hands := make([]hand, 0)
	s := bufio.NewScanner(file)

	cardScores := map[rune]int{
		'A': 14,
		'K': 13,
		'Q': 12,
		'J': 11,
		'T': 10,
		'9': 9,
		'8': 8,
		'7': 7,
		'6': 6,
		'5': 5,
		'4': 4,
		'3': 3,
		'2': 2,
	}

	day7 := d7{cardScores, hands}

	for s.Scan() {
		day7.hands = append(day7.hands, day7.buildHand(s.Text()))
	}

	sort.Sort(byWeakestHands(day7.hands))
	return &day7
}

func (d *d7) part1() int {
	score := 0

	for i, h := range d.hands {
		score += (i + 1) * h.bid
	}

	return score
}

func (d *d7) buildHand(s string) hand {
	parts := strings.Fields(s)

	cards := make([]int, 0)
	for _, c := range parts[0] {
		cards = append(cards, d.cardScores[c])
	}

	typ := scoreHand(cards)
	bid, _ := strconv.Atoi(parts[1])
	return hand{cards, typ, bid}
}

func scoreHand(cards []int) handType {
	cardToCount := make(map[int]int)

	for _, c := range cards {
		cardToCount[c] += 1
	}

	counts := make([]int, 0)
	for _, c := range cardToCount {
		counts = append(counts, c)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(counts)))

	if counts[0] == 1 {
		return highCard
	}

	if counts[0] == 2 {
		if counts[1] == 1 {
			return onePair
		}

		return twoPair
	}

	if counts[0] == 3 {
		if counts[1] == 1 {
			return threeOfAKind
		}

		return fullHouse
	}

	if counts[0] == 4 {
		return fourOfAKind
	}

	return fiveOfAKind
}

type byWeakestHands []hand

func (w byWeakestHands) Len() int      { return len(w) }
func (w byWeakestHands) Swap(i, j int) { w[i], w[j] = w[j], w[i] }
func (w byWeakestHands) Less(i, j int) bool {
	if w[i].typ != w[j].typ {
		return w[i].typ < w[j].typ
	}

	for l, val := range w[i].cards {
		if val != w[j].cards[l] {
			return val < w[j].cards[l]
		}
	}

	return false
}
