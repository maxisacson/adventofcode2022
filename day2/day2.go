package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	scoreMap := map[string]int{
		"X": 1, // rock
		"Y": 2, // paper
		"Z": 3, // scissor
	}

	beats := map[string]string{
		"A": "Y", // paper beats rock
		"B": "Z", // scissor beats paper
		"C": "X", // rock beats scissor
	}

	drawMap := map[string]string{
		"A": "X", // rock
		"B": "Y", // paper
		"C": "Z", // scissor
	}

	totalScore := 0
	for _, line := range lines {
		round := strings.Fields(line)
		opponent := round[0]
		player := round[1]

		score := scoreMap[player]
		if beats[opponent] == player {
			score += 6
		} else if drawMap[opponent] == player {
			score += 3
		}

		totalScore += score
	}

	if totalScore != 11666 && os.Args[1] == "input.txt" {
		panic("wrong answer!")
	}

	fmt.Println("--- Part 1 ---")
	fmt.Println("Total score:", totalScore)

	moveMap := map[string]int{
		"A": 0, // rock
		"B": 1, // paper
		"C": 2, // scisor
	}

	outcomeMap := map[string]int{
		"X": -1, // loose
		"Y": 0,  // draw
		"Z": 1,  // win
	}

	scoreList := []int{1, 2, 3}

	totalScore = 0
	for _, line := range lines {
		round := strings.Fields(line)
		opponent := moveMap[round[0]]
		outcome := outcomeMap[round[1]]

		player := (opponent + outcome + 3) % 3

		score := scoreList[player]
		if (opponent+1)%3 == player {
			score += 6 // win
		} else if opponent == player {
			score += 3 // draw
		}

		totalScore += score
	}

	if totalScore != 12767 && os.Args[1] == "input.txt" {
		panic("wrong answer!")
	}

	fmt.Println("--- Part 2 ---")
	fmt.Println("Total score:", totalScore)

}
