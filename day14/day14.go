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

type Board struct {
	Map      map[Vec]byte
	Min      Vec
	Max      Vec
	IsFinite bool
}

func (b *Board) DrawLine(start Vec, stop Vec) {
	if start.x == stop.x {
		y := start.y
		y1 := stop.y
		if y > y1 {
			y, y1 = y1, y
		}
		x := start.x
		for y <= y1 {
			b.DrawTileChecked(Vec{x, y}, '#')
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
			b.DrawTileChecked(Vec{x, y}, '#')
			x++
		}
	} else {
		panic("unexpected diagonal line")
	}
}

func (b *Board) DrawTileChecked(pos Vec, tile byte) {
	if pos.y > b.Max.y {
		b.Max.y = pos.y
	}
	if pos.y < b.Min.y {
		b.Min.y = pos.y
	}
	if pos.x > b.Max.x {
		b.Max.x = pos.x
	}
	if pos.x < b.Min.x {
		b.Min.x = pos.x
	}

	b.Map[pos] = tile
}

func MakeBoard(paths [][]Vec, source Vec) Board {

	board := Board{}
	board.Map = make(map[Vec]byte)
	board.Min = Vec{9999, 9999}
	board.Max = Vec{-9999, -9999}
	board.IsFinite = true

	for _, path := range paths {
		for i, point := range path[:len(path)-1] {
			board.DrawLine(point, path[i+1])
		}
	}
	board.DrawTileChecked(source, '+')

	return board
}

func (b *Board) IsAir(pos Vec) bool {
	_, ok := b.Map[pos]
	isAir := !ok
	isFloor := !b.IsFinite && pos.y == b.Max.y
	return isAir && !isFloor
}

func (b *Board) IsOutside(pos Vec) bool {
	inBoundsY := pos.y <= b.Max.y && pos.y >= b.Min.y
	inBoundsX := (pos.x <= b.Max.x && pos.x >= b.Min.x) || !b.IsFinite
	return !inBoundsX || !inBoundsY
}

func (b *Board) SimulateGrain(source Vec) bool {
	pos := source
	next := pos
	for true {
		next.y = pos.y + 1

		if b.IsOutside(next) {
			// sand is in the abyss
			return false
		}

		if b.IsAir(next) {
			pos = next
			continue
		}

		next.x = pos.x - 1
		next.y = pos.y + 1

		if b.IsOutside(next) {
			return false
		}

		if b.IsAir(next) {
			pos = next
			continue
		}

		next.x = pos.x + 1
		next.y = pos.y + 1

		if b.IsOutside(next) {
			return false
		}

		if b.IsAir(next) {
			pos = next
			continue
		}

		// sand comes to rest
		break
	}

	if b.IsFinite {
		b.Map[pos] = 'o'
	} else {
		b.DrawTileChecked(pos, 'o')
	}

	return true
}

func (b *Board) Draw() {

	for y := b.Min.y; y <= b.Max.y; y++ {
		for x := b.Min.x; x <= b.Max.x; x++ {
			if tile, ok := b.Map[Vec{x, y}]; ok {
				fmt.Print(string(tile))
			} else if !b.IsFinite && y == b.Max.y {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
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

	// Part 1
	board := MakeBoard(paths, source)

	count := 0
	for ; board.SimulateGrain(source); count++ {
	}

	// Part 1
	board = MakeBoard(paths, source)
	board.Max.y += 2
	board.IsFinite = false

	count2 := 0
	for ; board.Map[source] != 'o'; board.SimulateGrain(source) {
		count2++
	}
	return Result{count, count2}
}
