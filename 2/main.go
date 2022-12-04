package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func playToInt(moves []string) []int {
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

func playGame(p1, outcome int) int {
	var score = 0
	if outcome == 1 {
		score = lostRound(p1)
	} else if outcome == 2 {
		score = tieRound(p1)
	} else if outcome == 3 {
		score = wonRound(p1)
	}
	return score
}

func lostRound(p1 int) int {
	var score = p1 - 1
	if score <= 0 {
		score = 3
	}
	return score
}

func tieRound(p1 int) int {
	return p1 + 3
}

func wonRound(p1 int) int {
	var score = p1 + 1
	if score > 3 {
		score = 1
	}
	return score + 6
}

func main() {
	f, err := os.ReadFile("input.txt")
	check(err)

	var plays = strings.Split(string(f), "\n")
	var score = 0
	for i := 0; i < len(plays); i++ {
		var play = strings.Split(plays[i], " ")
		var playInt = playToInt(play)
		score += playGame(playInt[0], playInt[1])
	}
	fmt.Println("Score:", score)
}
