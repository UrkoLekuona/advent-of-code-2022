package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func main() {
	f, err := os.ReadFile("input.txt")
	check(err)

	var elfArray = strings.Split(string(f), "\n\n")
	var max = 0
	var maxIndex = -1
	for i := 0; i < len(elfArray); i++ {
		var current = countTotal(elfArray[i])
		if current > max {
			max = current
			maxIndex = i
		}
	}
	fmt.Println("The elf carrying the most calories is number", maxIndex+1, "with", max, "calories.")
}
