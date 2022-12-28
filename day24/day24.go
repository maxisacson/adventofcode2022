package day24

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

type Node struct {
	pos  Vec
	next []*Node
	prev *Node
	dist int
}

type Board struct {
	tiles    [][]Tile
	start    Vec
	goal     Vec
	path     *Node
	frontier []*Node
	ends     []*Node
}

func (b Board) String() string {
	nRows := len(b.tiles)
	offset := len(b.tiles[0]) + 1
	nCols := 2*len(b.tiles[0]) + 1

	bytes := make([][]byte, nRows)
	for i, row := range b.tiles {
		bytes[i] = make([]byte, nCols)
		for j, tile := range row {
			if tile.tile == 'B' {
				if len(tile.blizzards) == 1 {
					bytes[i][j] = tile.blizzards[0]
					bytes[i][j+offset] = tile.blizzards[0]
				} else {
					bytes[i][j] = byte('0' + len(tile.blizzards))
					bytes[i][j+offset] = byte('0' + len(tile.blizzards))
				}
			} else {
				bytes[i][j] = tile.tile
				bytes[i][j+offset] = tile.tile
			}
			// bytes[i][j] = tile.tile
			// bytes[i][j+offset] = tile.tile
			bytes[i][offset-1] = ' '
		}
	}

	for _, n := range b.frontier {
		bytes[n.pos.y][n.pos.x+offset] = '@'
	}

	lines := make([]string, nRows)
	for i := range lines {
		lines[i] = string(bytes[i])
	}

	return strings.Join(lines, "\n")
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

func (b *Board) IsBlizzard(x, y int) bool {
	return (b.tiles[y][x].tile == 'B')
}

func (b *Board) Adjacent(x, y int) []Vec {
	points := []Vec{}
	yMax := len(b.tiles) - 1
	xMax := len(b.tiles[0]) - 1

	for dy := -1; dy <= 1; dy++ {
		newY := y + dy
		if newY < 0 || newY > yMax {
			continue
		}
		for dx := -1; dx <= 1; dx++ {
			newX := x + dx
			if newX < 0 || newX > xMax {
				continue
			}

			if dx*dy != 0 {
				continue
			}

			if dx == 0 && dy == 0 {
				continue
			}

			if !b.IsWall(newX, newY) {
				points = append(points, Vec{newX, newY})
			}
		}
	}

	return points
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
					newX := x
					newY := y
					for !b.IsWall(newX-dx, newY-dy) {
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

	frontierMap := map[Vec]*Node{}
	for _, n := range b.frontier {
		for _, p := range b.Adjacent(n.pos.x, n.pos.y) {
			if n.prev != nil && p == n.prev.pos {
				continue
			}
			if b.IsBlizzard(p.x, p.y) {
				continue
			}
			next := new(Node)
			next.prev = n
			next.pos = p
			next.dist = n.dist + 1
			tmp, ok := frontierMap[p]
			if ok {
				if next.dist < tmp.dist {
					*tmp = *next
				}
			} else {
				frontierMap[p] = next
			}
			n.next = append(n.next, frontierMap[p])

			if p == b.goal {
				b.ends = append(b.ends, next)
			}
			// fmt.Println(p)
		}
		if !b.IsBlizzard(n.pos.x, n.pos.y) {
			next := new(Node)
			next.prev = n
			next.pos = n.pos
			next.dist = n.dist + 1
			n.next = append(n.next, next)
			frontierMap[next.pos] = next
		}
	}

	frontier := []*Node{}

	for _, v := range frontierMap {
		frontier = append(frontier, v)
	}

	b.frontier = frontier
}

func (b *Board) ShortestPath() []Vec {
	path := []Vec{}

	var next *Node
	minDist := -1

	for _, n := range b.ends {
		if minDist == -1 || n.dist < minDist {
			minDist = n.dist
			next = n
		}
	}

	for next != nil {
		path = append(path, next.pos)
		next = next.prev
	}

	return path
}

func Run(fileName string) Result {
	lines := utils.ReadFileToLines(fileName)

	nRows := len(lines)
	nCols := len(lines[0])

	board := Board{}
	board.tiles = make([][]Tile, nRows)
	tiles := make([][]Tile, nRows)
	for i, line := range lines {
		tiles[i] = make([]Tile, nCols)
		for j, b := range line {
			tiles[i][j] = MakeTile(byte(b))
			if i == 0 && b == '.' {
				board.start = Vec{j, i}
			}
			if i == nRows-1 && b == '.' {
				board.goal = Vec{j, i}
			}
		}
	}
	copy(board.tiles, tiles)
	root := Node{board.start, []*Node{}, nil, 0}
	board.path = &root
	board.frontier = append(board.frontier, board.path)

	// fmt.Println("start:", board.start)
	// fmt.Println("goal:", board.goal)
	// fmt.Println(board)
	// fmt.Println()

	// for i := 0; i < 18; i++ {
	for len(board.ends) == 0 {
		board.Round()
		// fmt.Println(len(board.frontier))
		// fmt.Println("Minute", i+1)
		// fmt.Println(board)
		// fmt.Println()
	}
	steps := len(board.ShortestPath()) - 1

	return Result{steps, 0}
}
