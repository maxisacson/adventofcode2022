package day22

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
	tiles   []string
	cursor  Cursor
	isCube  bool
	edgeMap map[Cursor]Cursor
}

func (b *Board) Move(move string) {
	switch move {
	case "L":
		b.RotateCursor(-1)
	case "R":
		b.RotateCursor(1)
	default:
		dist, _ := strconv.Atoi(move)
		b.MoveCursor(dist)
	}
}

func (b *Board) RotateCursor(dir int) {
	b.cursor.Rotate(dir)

	bytes := []byte(b.tiles[b.cursor.y])
	switch b.cursor.dir {
	case 0:
		bytes[b.cursor.x] = '>'
	case 1:
		bytes[b.cursor.x] = 'v'
	case 2:
		bytes[b.cursor.x] = '<'
	case 3:
		bytes[b.cursor.x] = '^'
	}
	b.tiles[b.cursor.y] = string(bytes)
}

func Direction(dir int) (int, int) {
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

	return dx, dy
}

func (b *Board) MoveCursor(steps int) {
	for step := 0; step < steps; step++ {
		dx, dy := Direction(b.cursor.dir)
		if dx != 0 && dy != 0 {
			panic("No diagonal steps!")
		}

		next := b.cursor
		if b.isCube {
			next = b.WrapXYCube(b.cursor.x+dx, b.cursor.y+dy, b.cursor.dir)
		} else {
			next.x = b.WrapXFlat(b.cursor.x+dx, b.cursor.y)
			next.y = b.WrapYFlat(b.cursor.x, b.cursor.y+dy)
		}

		if !b.IsValid(next.x, next.y) {
			break
		}
		b.cursor = next
		bytes := []byte(b.tiles[b.cursor.y])
		switch b.cursor.dir {
		case 0:
			bytes[b.cursor.x] = '>'
		case 1:
			bytes[b.cursor.x] = 'v'
		case 2:
			bytes[b.cursor.x] = '<'
		case 3:
			bytes[b.cursor.x] = '^'
		}
		b.tiles[b.cursor.y] = string(bytes)
	}

}

func (b *Board) IsValid(x, y int) bool {
	return b.tiles[y][x] != '#'
}

func (b *Board) WrapXYCube(x, y, dir int) Cursor {
	pos := Cursor{x, y, dir}
	dst, ok := b.edgeMap[pos]
	if ok {
		return dst
	}
	return pos
}

func (b *Board) WrapXFlat(x, y int) int {
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

func (b *Board) WrapYFlat(x, y int) int {
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
		panic(fmt.Sprintf("Take one step at a time! y = %d, (%d, %d)", y, yMin, yMax))
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

func MakeBoard(tiles []string, cursor Cursor, isCube bool) Board {
	board := Board{}
	board.tiles = make([]string, len(tiles))
	copy(board.tiles, tiles)
	board.cursor = cursor
	board.isCube = isCube

	bytes := []byte(board.tiles[cursor.y])
	switch cursor.dir {
	case 0:
		bytes[cursor.x] = '>'
	case 1:
		bytes[cursor.x] = 'v'
	case 2:
		bytes[cursor.x] = '<'
	case 3:
		bytes[cursor.x] = '^'
	}
	board.tiles[cursor.y] = string(bytes)

	if isCube {
		edgeMap := make(map[Cursor]Cursor)
		nRows := len(tiles)
		nCols := len(tiles[0])
		h := nRows / 3
		w := nCols / 4
		if h != w {
			panic(fmt.Sprintf("cube sides not equal: %d %d", w, h))
		}

		// 1 <-> 2
		for x := 0; x < w; x++ {
			// 1 -> 2
			src := Cursor{x + 2*w, -1, 3}
			dst := Cursor{w - 1 - x, h, 1}
			edgeMap[src] = dst

			// 2 -> 1
			src = Cursor{x, h - 1, 3}
			dst = Cursor{3*w - 1 - x, 0, 1}
			edgeMap[src] = dst
		}

		// 1 <-> 3
		for y := 0; y < h; y++ {
			// 1 -> 3
			src := Cursor{2*w - 1, y, 2}
			dst := Cursor{y + w, h, 1}
			edgeMap[src] = dst

			// 3 -> 1
			src = Cursor{y + w, h - 1, 3}
			dst = Cursor{2 * w, y, 0}
			edgeMap[src] = dst
		}

		// 1 <-> 6
		for y := 0; y < h; y++ {
			// 1 -> 6
			src := Cursor{3 * w, y, 0}
			dst := Cursor{4*w - 1, 3*h - 1 - y, 2}
			edgeMap[src] = dst

			// 6 -> 1
			src = Cursor{4 * w, 3*h - 1 - y, 0}
			dst = Cursor{3*w - 1, y, 2}
			edgeMap[src] = dst
		}

		// 2 <-> 5
		for x := 0; x < w; x++ {
			// 2 -> 5
			src := Cursor{x, 2 * h, 1}
			dst := Cursor{3*w - 1 - x, 3*h - 1, 3}
			edgeMap[src] = dst

			// 5 -> 2
			src = Cursor{3*w - 1 - x, 3 * h, 1}
			dst = Cursor{x, 2*h - 1, 3}
			edgeMap[src] = dst
		}

		// 2 <-> 6
		for y := 0; y < h; y++ {
			// 2 -> 6
			src := Cursor{-1, y + h, 2}
			dst := Cursor{4*w - 1 - y, 3*h - 1, 3}
			edgeMap[src] = dst

			// 6 -> 2
			src = Cursor{4*w - 1 - y, 3 * h, 1}
			dst = Cursor{0, y + h, 0}
			edgeMap[src] = dst
		}

		// 3 <-> 5
		for x := 0; x < w; x++ {
			// 3 -> 5
			src := Cursor{x + w, 2 * h, 1}
			dst := Cursor{2 * w, 3*h - 1 - x, 0}
			edgeMap[src] = dst

			// 5 -> 3
			src = Cursor{2*w - 1, 3*h - 1 - x, 2}
			dst = Cursor{x + w, 2*h - 1, 3}
			edgeMap[src] = dst
		}

		// 4 <-> 6
		for y := 0; y < h; y++ {
			// 4 -> 6
			src := Cursor{3 * w, y + h, 0}
			dst := Cursor{4*w - 1 - y, 2 * h, 1}
			edgeMap[src] = dst

			// 6 -> 4
			src = Cursor{4*w - 1 - y, 2*h - 1, 3}
			dst = Cursor{3*w - 1, y + h, 2}
			edgeMap[src] = dst
		}

		board.edgeMap = edgeMap
	}

	return board
}

func (b Board) String() string {
	return strings.Join(b.tiles, "\n")
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

	// Part 1
	board := MakeBoard(tiles, Cursor{x: x0, y: 0, dir: 0}, false)
	for _, move := range path {
		board.Move(move)
	}

	password := 1000*(board.cursor.y+1) + 4*(board.cursor.x+1) + board.cursor.dir

	// part 2
	board = MakeBoard(tiles, Cursor{x: x0, y: 0, dir: 0}, true)
	for _, move := range path {
		board.Move(move)
	}
	password2 := 1000*(board.cursor.y+1) + 4*(board.cursor.x+1) + board.cursor.dir

	return Result{password, password2}
}
