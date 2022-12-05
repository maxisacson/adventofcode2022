package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func contains(range1 []int, range2 []int) bool {
	// returns true if range1 fully contains range2
	return range1[0] <= range2[0] && range2[1] <= range1[1]
}

func overlaps(range1 []int, range2 []int) bool {
	// returns true if range1 and range2 overlap
	return !(range1[1] < range2[0] || range2[1] < range1[0])
}

func strings2ints(list []string) []int {
	var ret []int
	for _, s := range list {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		ret = append(ret, i)
	}
	return ret
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	countContains := 0
	countOverlaps := 0
	for _, line := range lines {
		pairs := strings.Split(line, ",")
		range1 := strings2ints(strings.Split(pairs[0], "-"))
		range2 := strings2ints(strings.Split(pairs[1], "-"))

		if contains(range1, range2) || contains(range2, range1) {
			countContains++
		}

		if overlaps(range1, range2) {
			countOverlaps++
		}
	}

	if countContains != 532 && os.Args[1] == "input.txt" {
		panic("wrong answer!")
	}

	fmt.Println("--- Part 1 ---")
	fmt.Println("Count:", countContains)

	if countOverlaps != 854 && os.Args[1] == "input.txt" {
		panic("wrong answer!")
	}

	fmt.Println("--- Part 2 ---")
	fmt.Println("Count:", countOverlaps)
}
