package main

import (
	"bufio"
	"fmt"
	"os"
)

func priority(char int) int {
	if 'a' <= char && char <= 'z' {
		return int(char - 'a' + 1)
	}

	return int(char - 'A' + 27)
}

func log2(x uint64) int {
	n := 0
	for x >>= 1; x != 0; x >>= 1 {
		n++
	}
	return n
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

	sum := 0
	for _, line := range lines {
		var count1 uint64
		var count2 uint64

		length := len(line)
		for i, c := range line {
			if i < length/2 {
				count1 |= (1 << (c - 'A'))
			} else {
				count2 |= (1 << (c - 'A'))
			}
		}

		result := count1 & count2
		char := log2(result) + 'A'
		sum += priority(char)
	}

	if sum != 7597 && os.Args[1] == "input.txt" {
		panic("wrong answer!")
	}

	fmt.Println("--- Part 1 ---")
	fmt.Println("Sum:", sum)

	sum = 0
	for i := 0; i < len(lines); i += 3 {
		var result uint64 = ^uint64(0)

		for j := 0; j < 3; j++ {
			var thisCount uint64
			for _, c := range lines[i+j] {
				thisCount |= (1 << (c - 'A'))
			}
			result &= thisCount
		}

		char := log2(result) + 'A'
		sum += priority(char)
	}

	if sum != 2607 && os.Args[1] == "input.txt" {
		panic("wrong answer!")
	}

	fmt.Println("--- Part 2 ---")
	fmt.Println("Sum:", sum)
}
