package main

import (
	"bufio"
	"fmt"
	"os"
)

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

	var insert struct{}

	fmt.Println("--- Part 1 ---")
	var result int
	for _, line := range lines {
		for i := 3; i < len(line); i++ {
			set := make(map[byte]struct{})
			for j := 3; j >= 0; j-- {
				set[line[i-j]] = insert
			}

			if len(set) == 4 {
				result = i + 1
				fmt.Println("Marker:", result, line[i-3:i+1])
				break
			}
		}
	}

	if result != 1929 && os.Args[1] == "input.txt" {
		panic("wrong answer!")
	}
}
