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
	handsP1 []hand
	handsP2 []hand
}

func newD7(path string) *d7 {
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	handsP1 := make([]hand, 0)
	handsP2 := make([]hand, 0)
	s := bufio.NewScanner(file)

	cardScoresP1 := map[rune]int{'A': 14, 'K': 13, 'Q': 12, 'J': 11, 'T': 10, '9': 9, '8': 8, '7': 7, '6': 6, '5': 5, '4': 4, '3': 3, '2': 2}
	cardScoresP2 := map[rune]int{'A': 14, 'K': 13, 'Q': 12, 'J': 1, 'T': 10, '9': 9, '8': 8, '7': 7, '6': 6, '5': 5, '4': 4, '3': 3, '2': 2}

	for s.Scan() {
		cardStr := s.Text()
		handsP1 = append(handsP1, buildHand(cardStr, cardScoresP1))
		handsP2 = append(handsP2, buildHand(cardStr, cardScoresP2))
	}

	sort.Sort(byWeakestHands(handsP1))
	sort.Sort(byWeakestHands(handsP2))

	return &d7{handsP1, handsP2}
}

func (d *d7) part1() int {
	score := 0

	for i, h := range d.handsP1 {
		score += (i + 1) * h.bid
	}

	return score
}

func (d *d7) part2() int {
	score := 0

	for i, h := range d.handsP2 {
		score += (i + 1) * h.bid
	}

	return score
}

func buildHand(s string, cardScores map[rune]int) hand {
	parts := strings.Fields(s)
	cards := make([]int, 0)

	for _, c := range parts[0] {
		cards = append(cards, cardScores[c])
	}

	var typ handType
	if cardScores['J'] == 11 {
		typ = scoreHandP1(cards)
	} else {
		typ = scoreHandP2(cards)
	}

	bid, _ := strconv.Atoi(parts[1])
	return hand{cards, typ, bid}
}

func scoreHandP1(cards []int) handType {
	cardToCount := make(map[int]int)

	for _, c := range cards {
		cardToCount[c] += 1
	}

	counts := make([]int, 0)
	for _, v := range cardToCount {
		counts = append(counts, v)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(counts)))

	mostCards := counts[0]
	secondmostCards := 0
	if len(counts) > 1 {
		secondmostCards += counts[1]
	}

	return scoreHandInner(mostCards, secondmostCards)
}

func scoreHandP2(cards []int) handType {
	cardToCount := make(map[int]int)

	for _, c := range cards {
		cardToCount[c] += 1
	}

	counts := make([]int, 0)
	for k, v := range cardToCount {
		if k == 1 {
			continue
		}
		counts = append(counts, v)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(counts)))

	mostCards := cardToCount[1]
	if len(counts) > 0 {
		mostCards += counts[0]
	}

	secondmostCards := 0
	if len(counts) > 1 {
		secondmostCards += counts[1]
	}

	return scoreHandInner(mostCards, secondmostCards)
}

func scoreHandInner(mostCards, secondmostCards int) handType {
	if mostCards == 5 {
		return fiveOfAKind
	}

	if mostCards == 4 {
		return fourOfAKind
	}

	if mostCards == 3 {
		if secondmostCards == 1 {
			return threeOfAKind
		}
		return fullHouse
	}

	if mostCards == 2 {
		if secondmostCards == 1 {
			return onePair
		}
		return twoPair
	}

	return highCard
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
