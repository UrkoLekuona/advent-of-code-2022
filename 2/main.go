package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type move int64

type play struct {
	name  string
	code  move
	score int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func movesToInt(moves []string) []int {
	var movesInt = make([]int, len(moves))
	for i := 0; i < len(moves); i++ {
		if (moves[i] == "A") || (moves[i] == "X") {
			movesInt[i] = 1
		} else if (moves[i] == "B") || (moves[i] == "Y") {
			movesInt[i] = 2
		} else if (moves[i] == "C") || (moves[i] == "Z") {
			movesInt[i] = 3
		} else {
			panic("Invalid move")
		}
	}
	return movesInt
}
func playGame(p1, p2 int) int {
	var diff = p1 - p2
	var winner int
	if diff < 0 {
		diff = -diff
	}
	if diff == 0 {
		return p2 + 3
	} else if diff == 1 {
		winner = int(math.Max(float64(p1), float64(p2)))
	} else if diff == 2 {
		winner = int(math.Min(float64(p1), float64(p2)))
	} else {
		panic("Invalid diff")
	}
	if winner == p2 {
		return p2 + 6
	} else {
		return p2
	}
}

func main() {
	f, err := os.ReadFile("input.txt")
	check(err)

	var plays = strings.Split(string(f), "\n")
	var score = 0
	for i := 0; i < len(plays); i++ {
		var moves = strings.Split(plays[i], " ")
		var movesInt = movesToInt(moves)
		score += playGame(movesInt[0], movesInt[1])
	}
	fmt.Println("Score:", score)
}
