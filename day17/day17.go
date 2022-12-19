package day17

import (
	"aoc22/utils"
	"fmt"
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

type Rock struct {
	pos    Vec
	blocks []Vec
	shape  string
}

func (r *Rock) Min() Vec {
	min := r.pos
	for _, b := range r.blocks {
		x := r.pos.x + b.x
		y := r.pos.y + b.y
		if x < min.x {
			min.x = x
		}
		if y < min.y {
			min.y = y
		}
	}
	return min
}

func (r *Rock) Max() Vec {
	max := r.pos
	for _, b := range r.blocks {
		x := r.pos.x + b.x
		y := r.pos.y + b.y
		if x > max.x {
			max.x = x
		}
		if y > max.y {
			max.y = y
		}
	}
	return max
}

func (r Rock) String() string {
	min := r.Min()
	max := r.Max()
	nRows := max.y - min.y + 1
	nCols := max.x - min.x + 1
	sprite := make([][]string, nRows)
	for y := 0; y < nRows; y++ {
		row := make([]string, nCols)
		for x := 0; x < nCols; x++ {
			row[x] = "."
		}
		sprite[y] = row
	}

	for _, b := range r.blocks {
		sprite[max.y-b.y][b.x] = "#"
	}

	var s string
	for _, row := range sprite {
		s += fmt.Sprintln(strings.Join(row, ""))
	}

	return s
}

func NewRock(shape rune) Rock {
	rock := Rock{}
	rock.shape = string(shape)
	switch shape {
	case '-':
		rock.blocks = make([]Vec, 4)
		rock.blocks[0] = Vec{0, 0}
		rock.blocks[1] = Vec{1, 0}
		rock.blocks[2] = Vec{2, 0}
		rock.blocks[3] = Vec{3, 0}
	case '+':
		rock.blocks = make([]Vec, 5)
		rock.blocks[0] = Vec{1, 0}
		rock.blocks[1] = Vec{0, 1}
		rock.blocks[2] = Vec{1, 1}
		rock.blocks[3] = Vec{2, 1}
		rock.blocks[4] = Vec{1, 2}
	case 'J':
		rock.blocks = make([]Vec, 5)
		rock.blocks[0] = Vec{0, 0}
		rock.blocks[1] = Vec{1, 0}
		rock.blocks[2] = Vec{2, 0}
		rock.blocks[3] = Vec{2, 1}
		rock.blocks[4] = Vec{2, 2}
	case 'I':
		rock.blocks = make([]Vec, 4)
		rock.blocks[0] = Vec{0, 0}
		rock.blocks[1] = Vec{0, 1}
		rock.blocks[2] = Vec{0, 2}
		rock.blocks[3] = Vec{0, 3}
	case 'o':
		rock.blocks = make([]Vec, 4)
		rock.blocks[0] = Vec{0, 0}
		rock.blocks[1] = Vec{1, 0}
		rock.blocks[2] = Vec{0, 1}
		rock.blocks[3] = Vec{1, 1}
	}

	return rock
}

func (r *Rock) Collides(o *Rock) bool {
	thisMin := r.Min()
	thisMax := r.Max()
	otherMin := o.Min()
	otherMax := o.Max()

	if thisMax.x < otherMin.x || thisMin.x > otherMax.x || thisMax.y < otherMin.y || thisMin.y > otherMax.y {
		return false
	}

	for _, b1 := range r.blocks {
		for _, b2 := range o.blocks {
			pos1 := Vec{r.pos.x + b1.x, r.pos.y + b1.y}
			pos2 := Vec{o.pos.x + b2.x, o.pos.y + b2.y}

			if pos1 == pos2 {
				return true
			}
		}
	}

	return false
}

type Board struct {
	rocks []Rock
	top   int
	width int
	jets  string
	j     int
}

func (b *Board) DropRock(shape rune) {
	rock := NewRock(shape)

	rock.pos.y = b.top + 3
	rock.pos.x = 2

	for true {
		b.MoveRockJet(&rock)
		if !b.MoveRockDown(&rock) {
			break
		}
	}

	testTop := rock.Max().y + 1
	if testTop > b.top {
		b.top = testTop
	}
	b.rocks = append(b.rocks, rock)
}

func (b *Board) MoveRockDown(rock *Rock) bool {
	rock.pos.y -= 1

	if rock.pos.y < 0 {
		rock.pos.y += 1
		return false
	}

	if b.CheckRockCollision(rock) {
		rock.pos.y += 1
		return false
	}

	return true
}

func (b *Board) MoveRockJet(rock *Rock) {
	jet := b.jets[b.j]
	b.j = (b.j + 1) % len(b.jets)
	switch jet {
	case '>':
		rock.pos.x += 1
		if !b.CheckPosition(rock) {
			rock.pos.x -= 1
		}
	case '<':
		rock.pos.x -= 1
		if !b.CheckPosition(rock) {
			rock.pos.x += 1
		}
	}
}

func (b *Board) CheckPosition(rock *Rock) bool {
	max := rock.Max()
	min := rock.Min()

	if max.x >= b.width || min.x < 0 || min.y < 0 {
		return false
	}

	return !b.CheckRockCollision(rock)
}

func (b *Board) CheckRockCollision(rock *Rock) bool {
	for _, other := range b.rocks {
		if rock.Collides(&other) {
			return true
		}
	}

	return false
}

func (b Board) String() string {
	nCols := b.width + 2
	nRows := b.top + 2

	sprite := make([][]string, nRows)
	for y := 0; y < nRows; y++ {
		row := make([]string, nCols)
		for x := 0; x < nCols; x++ {
			if y == 0 {
				if x == 0 || x == nCols-1 {
					row[x] = "+"
				} else {
					row[x] = "-"
				}
			} else if x == 0 {
				row[x] = "|"
			} else if x == nCols-1 {
				row[x] = fmt.Sprintf("| %d", y-1)
			} else {
				row[x] = " "
			}
		}
		sprite[y] = row
	}

	for _, rock := range b.rocks {
		for _, block := range rock.blocks {
			sprite[block.y+rock.pos.y+1][block.x+rock.pos.x+1] = rock.shape
		}
	}

	var s string
	for i := nRows - 1; i >= 0; i-- {
		s += fmt.Sprintln(strings.Join(sprite[i], ""))
	}

	return s
}

func Run(fileName string) Result {
	lines := utils.ReadFileToLines(fileName)

	jets := lines[0]

	rockOrder := "-+JIo"

	board := Board{}
	board.width = 7
	board.jets = jets
	for i := 0; i < 2022; i++ {
		shape := rockOrder[i%len(rockOrder)]
		board.DropRock(rune(shape))
	}

	return Result{board.top, 0}
}
