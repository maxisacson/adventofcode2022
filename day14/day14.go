package day14

import (
	"aoc22/utils"
	"fmt"
	"strconv"
	"strings"
)

type Result struct {
	part1 int
	part2 int
}

type Vec struct {
	x int
	y int
}

func ParsePath(line string) []Vec {
	path := []Vec{}

	fields := strings.Split(line, " -> ")
	for _, field := range fields {
		xy := strings.Split(field, ",")
		x, _ := strconv.Atoi(xy[0])
		y, _ := strconv.Atoi(xy[1])
		point := Vec{x, y}
		path = append(path, point)
	}

	return path
}

func DrawLine(board *[][]byte, xMin int, start Vec, stop Vec) {
	if start.x == stop.x {
		y := start.y
		y1 := stop.y
		if y > y1 {
			y, y1 = y1, y
		}
		x := start.x
		for y <= y1 {
			(*board)[y][x-xMin] = '#'
			y++
		}
	} else if start.y == stop.y {
		x := start.x
		x1 := stop.x
		if x > x1 {
			x, x1 = x1, x
		}
		y := start.y
		for x <= x1 {
			(*board)[y][x-xMin] = '#'
			x++
		}
	} else {
		panic("unexpected diagonal line")
	}
}

func MakeBoard(paths [][]Vec, source Vec) ([][]byte, int) {

	yMax := 0
	xMin := 99999
	xMax := -99999
	for _, path := range paths {
		for _, point := range path {
			if point.y > yMax {
				yMax = point.y
			}
			if point.x > xMax {
				xMax = point.x
			}
			if point.x < xMin {
				xMin = point.x
			}
		}
	}

	nRows := yMax + 1
	nCols := xMax - xMin + 1
	board := make([][]byte, nRows)
	for row := range board {
		board[row] = make([]byte, nCols)
		for col := range board[row] {
			board[row][col] = '.'
		}
	}

	for _, path := range paths {
		for i, point := range path[:len(path)-1] {
			DrawLine(&board, xMin, point, path[i+1])
		}
	}
	board[source.y][source.x-xMin] = '+'

	return board, xMin
}

func SimulateGrain(board *[][]byte, source Vec) bool {
	yMax := len(*board) - 1
	xMax := len((*board)[0]) - 1

	pos := source
	x, y := pos.x, pos.y
	for true {
		x = pos.x
		y = pos.y + 1

		if y > yMax {
			// sand is in the abyss
			return false
		}

		if (*board)[y][x] == '.' {
			pos = Vec{x, y}
			continue
		}

		x = pos.x - 1
		y = pos.y + 1
		if x < 0 || x > xMax {
			// sand is in the abyss
			return false
		}

		if (*board)[y][x] == '.' {
			pos = Vec{x, y}
			continue
		}

		x = pos.x + 1
		y = pos.y + 1
		if x < 0 || x > xMax {
			// sand is in the abyss
			return false
		}

		if (*board)[y][x] == '.' {
			pos = Vec{x, y}
			continue
		}

		// sand comes to rest
		break
	}

	(*board)[pos.y][pos.x] = 'o'

	return true
}

func DrawBoard(board *[][]byte) {
	for _, row := range *board {
		for _, col := range row {
			fmt.Print(string(col))
		}
		fmt.Println()
	}
	fmt.Println()
}

func Run(fileName string) Result {
	lines := utils.ReadFileToLines(fileName)

	paths := [][]Vec{}
	for _, line := range lines {
		path := ParsePath(line)
		paths = append(paths, path)
	}

	source := Vec{500, 0}
	board, xMin := MakeBoard(paths, source)
	DrawBoard(&board)
	source.x -= xMin

	count := 0
	for ; SimulateGrain(&board, source); count++ {
	}
	DrawBoard(&board)

	return Result{count, 0}
}
