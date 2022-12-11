package day1

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Result struct {
	part1 int
	part2 int
}

func Run(fileName string) Result {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var lines []string

	sum := 0
	var sums []int

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	for i, line := range lines {
		if line != "" {
			x, e := strconv.Atoi(line)
			if e != nil {
				panic(e)
			}
			sum += x
		}

		if line == "" || i == len(lines)-1 {
			sums = append(sums, sum)
			sum = 0
		}
	}

	sort.Ints(sums)
	maxSum := sums[len(sums)-1]
	topThree := sums[len(sums)-3:]
	topThreeSum := 0
	for _, s := range topThree {
		topThreeSum += s
	}

	if maxSum != 69795 && os.Args[1] == "input.txt" {
		panic("wrong answer!")
	}

	if topThreeSum != 208437 && os.Args[1] == "input.txt" {
		panic("wrong answer!")
	}

	fmt.Println("--- Part 1 ---")
	fmt.Println("Max sum:", maxSum)

	fmt.Println("--- Part 2 ---")
	fmt.Println("Top 3 sum:", topThreeSum)

	file.Close()

	return Result{maxSum, topThreeSum}
}
