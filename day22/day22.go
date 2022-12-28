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

const (
	Right int = iota
	Down      = iota
	Left      = iota
	Up        = iota
)

type Cursor struct {
	x   int
	y   int
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
}

func Direction(dir int) (int, int) {
	dx := 0
	dy := 0

	switch dir {
	case Right:
		dx = 1
	case Down:
		dy = 1
	case Left:
		dx = -1
	case Up:
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
		panic(fmt.Sprintf("Take one step at a time! y = %d, (%d, %d)", x, xMin, xMax))
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

func GetOpposite(dir int) int {
	return (dir + 2) % 4
}

func MakeEdgeMap(tiles []string, sides [6]CubeSide) map[Cursor]Cursor {
	edgeMap := make(map[Cursor]Cursor)

	for _, side := range sides {
		for dir, link := range side.links {
			if link == 0 {
				continue
			}

			j := link - 1
			linkSide := side.linkSide[dir]
			linkParity := side.linkParity[dir]
			otherSide := sides[j]
			if link != otherSide.label {
				panic("mismatched link!")
			}

			opDir := GetOpposite(linkSide)
			points := side.GetPoints(dir, 1)
			otherPoints := otherSide.GetPoints(linkSide, linkParity)

			switch dir {
			case Right:
				for i := range points {
					points[i].x += 1
				}
			case Down:
				for i := range points {
					points[i].y += 1
				}
			case Left:
				for i := range points {
					points[i].x -= 1
				}
			case Up:
				for i := range points {
					points[i].y -= 1
				}
			default:
				panic(fmt.Sprintf("unknown side: %v", dir))
			}

			for i, p := range points {
				src := Cursor{p.x, p.y, dir}
				dst := Cursor{otherPoints[i].x, otherPoints[i].y, opDir}
				edgeMap[src] = dst
			}
		}
	}

	return edgeMap
}

type CubeSide struct {
	x int
	y int
	w int
	h int

	label      int
	links      [4]int
	linkSide   [4]int
	linkParity [4]int
}

type Vec struct {
	x int
	y int
}

func (s *CubeSide) GetPoints(side, parity int) []Vec {
	points := []Vec{}

	switch side {
	case Right:
		x := s.x + s.w - 1
		for i := 0; i < s.h; i++ {
			y := s.y + i
			points = append(points, Vec{x, y})
		}
	case Down:
		y := s.y + s.h - 1
		for i := 0; i < s.w; i++ {
			x := s.x + i
			points = append(points, Vec{x, y})
		}
	case Left:
		x := s.x
		for i := 0; i < s.h; i++ {
			y := s.y + i
			points = append(points, Vec{x, y})
		}
	case Up:
		y := s.y
		for i := 0; i < s.w; i++ {
			x := s.x + i
			points = append(points, Vec{x, y})
		}
	default:
		panic(fmt.Sprintf("unknown side: %v", side))
	}

	if parity < 0 {
		N := len(points)
		tmp := make([]Vec, N)
		for i, x := range points {
			tmp[N-1-i] = x
		}
		points = tmp
	}
	return points
}

func MakeBoard(tiles []string, cursor Cursor, isCube bool, cubeSides [6]CubeSide) Board {
	board := Board{}
	board.tiles = make([]string, len(tiles))
	copy(board.tiles, tiles)
	board.cursor = cursor
	board.isCube = isCube

	if board.isCube {
		board.edgeMap = MakeEdgeMap(tiles, cubeSides)
	}

	return board
}

func (b Board) String() string {
	return strings.Join(b.tiles, "\n")
}

func MakeCubeSides(w, h int) [6]CubeSide {
	cubeSides := [6]CubeSide{}
	if w == 4 {
		cubeSides[0] = CubeSide{
			x: 2 * w,
			y: 0,
			w: w,
			h: h,

			label:      1,
			links:      [4]int{6, 0, 3, 2},
			linkSide:   [4]int{Right, 0, Up, Up},
			linkParity: [4]int{-1, 0, 1, -1},
		}
		cubeSides[1] = CubeSide{
			x: 0,
			y: h,
			w: w,
			h: h,

			label:      2,
			links:      [4]int{0, 5, 6, 1},
			linkSide:   [4]int{0, Down, Down, Up},
			linkParity: [4]int{0, -1, -1, -1},
		}
		cubeSides[2] = CubeSide{
			x: w,
			y: h,
			w: w,
			h: h,

			label:      3,
			links:      [4]int{0, 5, 0, 1},
			linkSide:   [4]int{0, Left, 0, Left},
			linkParity: [4]int{0, -1, 0, 1},
		}
		cubeSides[3] = CubeSide{
			x: 2 * w,
			y: h,
			w: w,
			h: h,

			label:      4,
			links:      [4]int{6, 0, 0, 0},
			linkSide:   [4]int{Up, 0, 0, 0},
			linkParity: [4]int{-1, 0, 0, 0},
		}
		cubeSides[4] = CubeSide{
			x: 2 * w,
			y: 2 * h,
			w: w,
			h: h,

			label:      5,
			links:      [4]int{0, 2, 3, 0},
			linkSide:   [4]int{0, Down, Down, 0},
			linkParity: [4]int{0, -1, -1, 0},
		}
		cubeSides[5] = CubeSide{
			x: 3 * w,
			y: 2 * h,
			w: w,
			h: h,

			label:      6,
			links:      [4]int{1, 2, 0, 4},
			linkSide:   [4]int{Right, Left, 0, Right},
			linkParity: [4]int{-1, -1, 0, -1},
		}
	} else if w == 50 {
		cubeSides[0] = CubeSide{
			x: w,
			y: 0,
			w: w,
			h: h,

			label:      1,
			links:      [4]int{0, 0, 4, 6},
			linkSide:   [4]int{0, 0, Left, Left},
			linkParity: [4]int{0, 0, -1, 1},
		}
		cubeSides[1] = CubeSide{
			x: 2 * w,
			y: 0,
			w: w,
			h: h,

			label:      2,
			links:      [4]int{5, 3, 0, 6},
			linkSide:   [4]int{Right, Right, 0, Down},
			linkParity: [4]int{-1, 1, 0, 1},
		}
		cubeSides[2] = CubeSide{
			x: w,
			y: h,
			w: w,
			h: h,

			label:      3,
			links:      [4]int{2, 0, 4, 0},
			linkSide:   [4]int{Down, 0, Up, 0},
			linkParity: [4]int{1, 0, 1, 0},
		}
		cubeSides[3] = CubeSide{
			x: 0,
			y: 2 * h,
			w: w,
			h: h,

			label:      4,
			links:      [4]int{0, 0, 1, 3},
			linkSide:   [4]int{0, 0, Left, Left},
			linkParity: [4]int{0, 0, -1, 1},
		}
		cubeSides[4] = CubeSide{
			x: w,
			y: 2 * h,
			w: w,
			h: h,

			label:      5,
			links:      [4]int{2, 6, 0, 0},
			linkSide:   [4]int{Right, Right, 0, 0},
			linkParity: [4]int{-1, 1, 0, 0},
		}
		cubeSides[5] = CubeSide{
			x: 0,
			y: 3 * h,
			w: w,
			h: h,

			label:      6,
			links:      [4]int{5, 2, 1, 0},
			linkSide:   [4]int{Down, Up, Up, 0},
			linkParity: [4]int{1, 1, 1, 0},
		}
	}

	return cubeSides
}

func Run(fileName string, sideSize int) Result {
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
	board := MakeBoard(tiles, Cursor{x: x0, y: 0, dir: Right}, false, [6]CubeSide{})
	for _, move := range path {
		board.Move(move)
	}

	password := 1000*(board.cursor.y+1) + 4*(board.cursor.x+1) + board.cursor.dir

	// part 2
	cubeSides := MakeCubeSides(sideSize, sideSize)
	board = MakeBoard(tiles, Cursor{x: x0, y: 0, dir: Right}, true, cubeSides)
	for _, move := range path {
		board.Move(move)
	}
	password2 := 1000*(board.cursor.y+1) + 4*(board.cursor.x+1) + board.cursor.dir

	return Result{password, password2}
}
