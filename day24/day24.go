package day24

import (
	"aoc22/utils"
	"fmt"
)

type Result struct {
	part1 int
	part2 int
}

type Vec struct {
	x int
	y int
}

type Tile struct {
	tile      byte
	blizzards []byte
}

func MakeTile(b byte) Tile {
	tile := Tile{}
	tile.tile = b

	switch b {
	case '>':
		fallthrough
	case 'v':
		fallthrough
	case '<':
		fallthrough
	case '^':
		tile.tile = 'B'
		tile.blizzards = append(tile.blizzards, b)
	}

	return tile
}

func Pop(list *[]byte) byte {
	n := len(*list) - 1
	r := (*list)[n]
	*list = (*list)[:n]
	return r
}

type Board struct {
	tiles [][]Tile
	start Vec
	goal  Vec
}

func (b Board) String() string {
	s := ""
	for i, row := range b.tiles {
		for _, tile := range row {
			if tile.tile == 'B' {
				if len(tile.blizzards) == 1 {
					s += string(tile.blizzards[0])
				} else {
					s += fmt.Sprint(len(tile.blizzards))
				}
			} else {
				s += fmt.Sprint(string(tile.tile))
			}
		}
		if i < len(b.tiles)-1 {
			s += "\n"
		}
	}
	return s
}

func GetBlizzardDir(b byte) (int, int) {
	switch b {
	case '>':
		return 1, 0
	case 'v':
		return 0, 1
	case '<':
		return -1, 0
	case '^':
		return 0, -1
	}

	return 0, 0
}

func (b *Board) IsWall(x, y int) bool {
	return (b.tiles[y][x].tile == '#')
}

func (b *Board) Round() {
	tiles := make([][]Tile, len(b.tiles))
	for i := range b.tiles {
		tiles[i] = make([]Tile, len(b.tiles[i]))
	}

	// create empty board
	for y, row := range b.tiles {
		for x, tile := range row {
			if tile.tile == '#' {
				tiles[y][x].tile = '#'
			} else {
				tiles[y][x].tile = '.'
			}
		}
	}

	// simulate blizzards
	for y, row := range b.tiles {
		for x, tile := range row {
			if tile.tile != 'B' {
				continue
			}
			for _, bliz := range tile.blizzards {
				dx, dy := GetBlizzardDir(bliz)
				if b.IsWall(x+dx, y+dy) {
					// Create new blizzard at opposite wall
					newX := x - dx
					newY := y - dy
					for !b.IsWall(newX, newY) {
						newX -= dx
						newY -= dy
					}
					tiles[newY][newX].tile = 'B'
					tiles[newY][newX].blizzards = append(tiles[newY][newX].blizzards, bliz)
				} else {
					newX := x + dx
					newY := y + dy
					tiles[newY][newX].tile = 'B'
					tiles[newY][newX].blizzards = append(tiles[newY][newX].blizzards, bliz)
				}
			}
		}
	}

	b.tiles = tiles
}

func Run(fileName string) Result {
	lines := utils.ReadFileToLines(fileName)

	nRows := len(lines)
	nCols := len(lines[0])

	board := Board{}
	board.tiles = make([][]Tile, nRows)
	for i, line := range lines {
		board.tiles[i] = make([]Tile, nCols)
		for j, b := range line {
			board.tiles[i][j] = MakeTile(byte(b))
			if i == 0 && b == '.' {
				board.start = Vec{j, i}
			}
			if i == nRows-1 && b == '.' {
				board.goal = Vec{j, i}
			}
		}
	}

	fmt.Println("start:", board.start)
	fmt.Println("goal:", board.goal)
	fmt.Println(board)
	fmt.Println()

	for i := 0; i < 10; i++ {
		board.Round()
		fmt.Println(board)
		fmt.Println()
	}

	return Result{len(lines), len(lines)}
}
