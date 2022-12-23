package day23

import (
	"aoc22/utils"
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

type Elf struct {
	pos  Vec
	next Vec
}

func (e *Elf) GetNext(dir int) Vec {
	switch dir {
	case North:
		return Vec{e.pos.x, e.pos.y - 1}
	case South:
		return Vec{e.pos.x, e.pos.y + 1}
	case West:
		return Vec{e.pos.x - 1, e.pos.y}
	case East:
		return Vec{e.pos.x + 1, e.pos.y}
	}

	return e.pos
}

type Board struct {
	tiles [][]byte
	min   Vec
	elfs  []Elf
	dirs  []int
}

const (
	North int = iota
	South
	West
	East
)

func (b Board) String() string {
	s := ""
	for i, row := range b.tiles {
		s += string(row)
		if i < len(b.tiles)-1 {
			s += "\n"
		}
	}

	return s
}

func (b *Board) IsEmpty(x, y int) bool {
	row, col := b.GetRowCol(x, y)
	if 0 <= row && row < len(b.tiles) && 0 <= col && col < len(b.tiles[0]) {
		return b.tiles[row][col] == '.'
	}

	return true
}

func (b *Board) ShouldMove(elf *Elf) bool {
	for dy := -1; dy <= 1; dy++ {
		y := elf.pos.y + dy
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			x := elf.pos.x + dx
			if !b.IsEmpty(x, y) {
				return true
			}
		}
	}
	return false
}

func (b *Board) IsValid(dir int, pos Vec) bool {
	switch dir {
	case North:
		dy := -1
		for dx := -1; dx <= 1; dx++ {
			if !b.IsEmpty(pos.x+dx, pos.y+dy) {
				return false
			}
		}
	case South:
		dy := 1
		for dx := -1; dx <= 1; dx++ {
			if !b.IsEmpty(pos.x+dx, pos.y+dy) {
				return false
			}
		}
	case West:
		dx := -1
		for dy := -1; dy <= 1; dy++ {
			if !b.IsEmpty(pos.x+dx, pos.y+dy) {
				return false
			}
		}
	case East:
		dx := 1
		for dy := -1; dy <= 1; dy++ {
			if !b.IsEmpty(pos.x+dx, pos.y+dy) {
				return false
			}
		}
	}

	return true
}

func (b *Board) Round() bool {
	// first half
	moves := make(map[Vec]int)
	anyMoves := false
	for i := range b.elfs {
		elf := &b.elfs[i]
		elf.next = elf.pos
		if !b.ShouldMove(elf) {
			continue
		}
		for _, dir := range b.dirs {
			if b.IsValid(dir, elf.pos) {
				elf.next = elf.GetNext(dir)
				break
			}
		}
		if elf.next == elf.pos {
			continue
		}
		anyMoves = true

		_, ok := moves[elf.next]
		if ok {
			moves[elf.next] += 1
		} else {
			moves[elf.next] = 1
		}
	}

	if !anyMoves {
		return false
	}

	// second half
	for i := range b.elfs {
		elf := &b.elfs[i]
		if elf.next == elf.pos {
			continue
		}

		if moves[elf.next] == 1 {
			b.MoveElf(elf)
		}
	}

	b.dirs = append(b.dirs[1:], b.dirs[0])

	return true
}

func (b *Board) GetRowCol(x, y int) (int, int) {
	row := y - b.min.y
	col := x - b.min.x
	return row, col
}

func (b *Board) MoveElf(elf *Elf) {
	row, col := b.GetRowCol(elf.pos.x, elf.pos.y)
	b.tiles[row][col] = '.'

	elf.pos = elf.next
	row, col = b.GetRowCol(elf.pos.x, elf.pos.y)

	if row < 0 {
		nCols := len(b.tiles[0])
		newRow := []byte(strings.Repeat(".", nCols))
		b.tiles = append([][]byte{newRow}, b.tiles...)
		b.min.y -= 1
	}
	if row >= len(b.tiles) {
		nCols := len(b.tiles[0])
		newRow := []byte(strings.Repeat(".", nCols))
		b.tiles = append(b.tiles, newRow)
	}
	if col < 0 {
		for i := range b.tiles {
			b.tiles[i] = append([]byte{'.'}, b.tiles[i]...)
		}
		b.min.x -= 1
	}
	if col >= len(b.tiles[0]) {
		for i := range b.tiles {
			b.tiles[i] = append(b.tiles[i], '.')
		}
	}
	row, col = b.GetRowCol(elf.pos.x, elf.pos.y)

	b.tiles[row][col] = '#'
}

func Run(fileName string) Result {
	lines := utils.ReadFileToLines(fileName)

	nRows := len(lines)
	nCols := len(lines[0])

	tiles := make([][]byte, nRows)
	elfs := []Elf{}
	dirs := []int{North, South, West, East}
	for i, line := range lines {
		tiles[i] = make([]byte, nCols)
		for j, b := range line {
			tiles[i][j] = byte(b)
			if b == '#' {
				elfs = append(elfs, Elf{Vec{j, i}, Vec{j, i}})
			}
		}
	}

	board := Board{}
	board.tiles = tiles
	board.elfs = elfs
	board.dirs = dirs

	count := 0
	i := 1
	for ; board.Round(); i++ {
		if i == 10 {
			for _, row := range board.tiles {
				for _, tile := range row {
					if tile == '.' {
						count++
					}
				}
			}
		}
	}

	return Result{count, i}
}
