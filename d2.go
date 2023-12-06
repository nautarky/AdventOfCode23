package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const d2InputPath = "input/2.txt"

type game map[string]int

func d2() (int, int) {
	file, err := os.Open(d2InputPath)
	check(err)
	defer file.Close()

	redLimit, greenLimit, blueLimit := 12, 13, 14
	p1, p2, gameId := 0, 0, 1

	s := bufio.NewScanner(file)
	for s.Scan() {
		line := s.Text()
		g := parseGame(line)

		if g["red"] <= redLimit && g["green"] <= greenLimit && g["blue"] <= blueLimit {
			p1 += gameId
		}

		p2 += g["red"] * g["green"] * g["blue"]
		gameId++
	}

	return p1, p2
}

func parseGame(g string) game {
	minViableGame := make(game)

	gameParts := strings.Fields(g)
	for i := 2; i < len(gameParts); i += 2 {
		numCubes, _ := strconv.Atoi(gameParts[i])
		color := strings.Trim(gameParts[i+1], ",;")
		minViableGame[color] = max(numCubes, minViableGame[color])
	}

	return minViableGame
}
