package day8

import (
	"aoc22/utils"
)

type Result struct {
	part1 int
	part2 int
}

func Run(fileName string) Result {
	lines := utils.ReadFileToLines(fileName)

	nRows := len(lines)
	nCols := len(lines[0])

	isVisibleLeft := make([][]int, nRows)
	isVisibleRight := make([][]int, nRows)
	isVisibleTop := make([][]int, nRows)
	isVisibleBottom := make([][]int, nRows)

	heights := make([][]int, nRows)

	for i, line := range lines {
		heights[i] = make([]int, nCols)
		for j, c := range line {
			heights[i][j] = int(c - '0')
		}
		isVisibleLeft[i] = make([]int, nCols)
		isVisibleRight[i] = make([]int, nCols)
		isVisibleTop[i] = make([]int, nCols)
		isVisibleBottom[i] = make([]int, nCols)
	}

	for i := 1; i < nRows-1; i++ {
		for j := 1; j < nCols-1; j++ {

			h := heights[i][j]
			if h > heights[i][0] {
				heights[i][0] = h
				isVisibleLeft[i][j] = 1
			}

			if h > heights[0][j] {
				heights[0][j] = h
				isVisibleTop[i][j] = 1
			}

			h = heights[i][nCols-j-1]
			if h > heights[i][nCols-1] {
				heights[i][nCols-1] = h
				isVisibleRight[i][nCols-j-1] = 1
			}

			h = heights[nRows-i-1][j]
			if h > heights[nRows-1][j] {
				heights[nRows-1][j] = h
				isVisibleBottom[nRows-i-1][j] = 1
			}
		}
	}

	nVisible := 2*nRows + 2*(nCols-2)

	for i, line := range lines {
		for j := range line {
			isVisible := isVisibleLeft[i][j] | isVisibleRight[i][j] | isVisibleTop[i][j] | isVisibleBottom[i][j]
			nVisible += isVisible
		}
	}

	return Result{nVisible, 0}
}
