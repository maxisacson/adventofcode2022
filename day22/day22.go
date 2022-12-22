package day22

import (
	"aoc22/utils"
	"strconv"
	"strings"
)

type Result struct {
	part1 int
	part2 int
}

type Cursor struct {
	x int
	y int

	// 0: right >
	// 1: down  v
	// 2: left  <
	// 3: up    ^
	dir int
}

func (c *Cursor) Rotate(dir int) {
	c.dir += dir
	c.dir += 4
	c.dir %= 4
}

type Board struct {
	tiles  []string
	cursor Cursor
}

func (b *Board) Move(move string) {
	switch move {
	case "L":
		b.cursor.Rotate(-1)
	case "R":
		b.cursor.Rotate(1)
	default:
		dist, _ := strconv.Atoi(move)
		b.MoveCursor(dist)
	}
}

func (b *Board) MoveCursor(steps int) {
	dir := b.cursor.dir

	dx := 0
	dy := 0

	switch dir {
	case 0: //right
		dx = 1
	case 1: // down
		dy = 1
	case 2: // left
		dx = -1
	case 3: // up
		dy = -1
	}

	for step := 0; step < steps; step++ {
		x := b.WrapX(b.cursor.x+dx, b.cursor.y)
		y := b.WrapY(b.cursor.x, b.cursor.y+dy)
		if !b.IsValid(x, y) {
			break
		}
		b.cursor.x = x
		b.cursor.y = y
	}
}

func (b *Board) IsValid(x, y int) bool {
	return b.tiles[y][x] != '#'
}

func (b *Board) WrapX(x, y int) int {
	xMin := 0
	xMax := len(b.tiles[0]) - 1

	row := b.tiles[y]

	for j, c := range row {
		if c != ' ' {
			xMin = j
			break
		}
	}

	for j := xMax; j >= xMin; j-- {
		if row[j] != ' ' {
			xMax = j
			break
		}
	}

	if (xMin-x) > 1 || (x-xMax) > 1 {
		panic("Take one step at a time!")
	}

	newX := x
	if x < xMin {
		newX = xMax
	}

	if x > xMax {
		newX = xMin
	}

	return newX
}

func (b *Board) WrapY(x, y int) int {
	yMin := 0
	yMax := len(b.tiles) - 1

	for i, row := range b.tiles {
		if row[x] != ' ' {
			yMin = i
			break
		}
	}

	for i := yMax; i >= yMin; i-- {
		if b.tiles[i][x] != ' ' {
			yMax = i
			break
		}
	}

	if (yMin-y) > 1 || (y-yMax) > 1 {
		panic("Take one step at a time!")
	}

	newY := y
	if y < yMin {
		newY = yMax
	}

	if y > yMax {
		newY = yMin
	}

	return newY
}

func Run(fileName string) Result {
	lines := utils.ReadFileToLines(fileName)

	nRows := len(lines) - 2
	nCols := 0
	tiles := make([]string, nRows)
	for i, line := range lines {
		if line == "" {
			break
		}
		tiles[i] = line
		if len(line) > nCols {
			nCols = len(line)
		}
	}
	x0 := 0
	for x, c := range tiles[0] {
		if c == '.' {
			x0 = x
			break
		}
	}
	for i, line := range tiles {
		if len(line) < nCols {
			tiles[i] += strings.Repeat(" ", nCols-len(line))
		}
	}

	board := Board{tiles: tiles, cursor: Cursor{x: x0, y: 0, dir: 0}}

	pathStr := lines[nRows+1]
	path := []string{}

	i := 0
	current := ""
	for i < len(pathStr) {
		if pathStr[i] == 'R' || pathStr[i] == 'L' {
			if current != "" {
				path = append(path, current)
			}
			path = append(path, string(pathStr[i]))
			current = ""
			i++
		}
		current += string(pathStr[i])
		i++
	}
	path = append(path, current)

	for _, move := range path {
		board.Move(move)
	}

	password := 1000*(board.cursor.y+1) + 4*(board.cursor.x+1) + board.cursor.dir

	return Result{password, 0}
}
