package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type elf struct {
	index int
	total int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func countTotal(elementString string) int {
	var elementArray = strings.Split(elementString, "\n")
	var total = 0
	for i := 0; i < len(elementArray); i++ {
		elementInt, err := strconv.Atoi(elementArray[i])
		check(err)
		total += elementInt
	}
	return total
}

func minOfThree(topThree []elf) (elf, int) {
	var min = topThree[0]
	var minIndex = 0
	for i := 1; i < len(topThree); i++ {
		if topThree[i].total < min.total {
			min = topThree[i]
			minIndex = i
		}
	}
	return min, minIndex
}

func main() {
	f, err := os.ReadFile("input.txt")
	check(err)

	var elfArray = strings.Split(string(f), "\n\n")
	var smallestOfTopThree elf = elf{-1, 0}
	var topThree = make([]elf, 3)
	var smallestOfTopThreeIndex = 0
	for i := 0; i < len(elfArray); i++ {
		var current = countTotal(elfArray[i])
		if current > smallestOfTopThree.total {
			topThree[smallestOfTopThreeIndex] = elf{i, current}
			smallestOfTopThree, smallestOfTopThreeIndex = minOfThree(topThree)
		}
	}
	var totalCalories = topThree[0].total + topThree[1].total + topThree[2].total
	fmt.Println("The top three elfs carrying the most calories are numbers", topThree[0].index+1, ",", topThree[1].index+1, "and", topThree[2].index+1, "with a total of", totalCalories, "calories.")
}
