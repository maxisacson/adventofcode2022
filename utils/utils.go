package utils

import (
	"bufio"
	"os"
)

func ReadFileToLines(fileName string) []string {
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

	return lines
}
