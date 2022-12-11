package day6

import (
	"bufio"
	"fmt"
	"os"
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
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	var insert struct{}

	fmt.Println("--- Part 1 ---")
	var startOfPacket int

	for _, line := range lines {
		startOfPacket = 0
		for i := 3; i < len(line); i++ {
			markerSet := make(map[byte]struct{})

			for j := 3; j >= 0; j-- {
				markerSet[line[i-j]] = insert
			}

			if startOfPacket == 0 && len(markerSet) == 4 {
				startOfPacket = i + 1
				fmt.Println("Marker:", startOfPacket, line[i-3:i+1])
			}
		}
	}

	if startOfPacket != 1929 && os.Args[1] == "input.txt" {
		panic("wrong answer!")
	}

	fmt.Println("--- Part 2 ---")
	var startOfMessage int

	for _, line := range lines {
		startOfMessage = 0
		for i := 13; i < len(line); i++ {
			messageSet := make(map[byte]struct{})

			for j := 13; j >= 0; j-- {
				messageSet[line[i-j]] = insert
			}

			if startOfMessage == 0 && len(messageSet) == 14 {
				startOfMessage = i + 1
				fmt.Println("Message:", startOfMessage, line[i-13:i+1])
			}
		}
	}

	if startOfMessage != 3298 && os.Args[1] == "input.txt" {
		panic("wrong answer!")
	}

	return Result{startOfPacket, startOfMessage}
}
