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

	// Part 2
	bestScore := 0
	for i, row := range heights {
		for j, h := range row {
			// look up
			score := 0
			distance := 0
			for k := i - 1; k >= 0; k-- {
				distance++
				if heights[k][j] >= h {
					break
				}
			}
			score = distance
			// look down
			distance = 0
			for k := i + 1; k < nRows; k++ {
				distance++
				if heights[k][j] >= h {
					break
				}
			}
			score *= distance

			// look left
			distance = 0
			for k := j - 1; k >= 0; k-- {
				distance++
				if heights[i][k] >= h {
					break
				}
			}
			score *= distance

			// look right
			distance = 0
			for k := j + 1; k < nCols; k++ {
				distance++
				if heights[i][k] >= h {
					break
				}
			}
			score *= distance

			if score > bestScore {
				bestScore = score
			}
		}
	}

	// Part 1
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

	return Result{nVisible, bestScore}
}
