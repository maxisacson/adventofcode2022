package day9

import (
	"aoc22/utils"
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

type Rope struct {
	head Vec
	tail Vec
}

func sign(x int) int {
	if x < 0 {
		return -1
	}
	return 1
}

func abs(x int) int {
	return x * sign(x)
}

func norm(x int) int {
	if x == 0 {
		return 0
	}

	return sign(x)
}

func (r *Rope) Move(dx int, dy int) {
	r.head.x += dx
	r.head.y += dy

	r.MoveTail()
}

func (r *Rope) MoveTail() {
	dx := r.head.x - r.tail.x
	dy := r.head.y - r.tail.y

	if abs(dx) <= 1 && abs(dy) <= 1 {
		return
	}

	// This works only if head moves one square at a time
	r.tail.x += norm(dx)
	r.tail.y += norm(dy)
}

type LongRope struct {
	head Vec
	tail Vec
	body []Vec
}

func NewLongRope(size int) LongRope {
	return LongRope{
		body: make([]Vec, size),
	}
}

func (r *LongRope) Move(dx int, dy int) {
	r.body[0].x += dx
	r.body[0].y += dy

	r.MoveTail()
}

func (r *LongRope) MoveNextTowards(index int) {
	dx := r.body[index].x - r.body[index+1].x
	dy := r.body[index].y - r.body[index+1].y

	if abs(dx) <= 1 && abs(dy) <= 1 {
		return
	}

	// This works only if head moves one square at a time
	r.body[index+1].x += norm(dx)
	r.body[index+1].y += norm(dy)
}

func (r *LongRope) MoveTail() {
	for i := 1; i < len(r.body); i++ {
		r.MoveNextTowards(i - 1)
	}
}

func Run(fileName string) Result {
	lines := utils.ReadFileToLines(fileName)

	var visit struct{}

	visited := make(map[Vec]struct{})
	longRopeVisited := make(map[Vec]struct{})

	rope := Rope{}
	longRope := NewLongRope(10)
	for _, line := range lines {
		fields := strings.Fields(line)
		dir := fields[0]
		count, _ := strconv.Atoi(fields[1])

		dx := 0
		dy := 0
		switch dir {
		case "D":
			dy = 1
		case "U":
			dy = -1
		case "R":
			dx = 1
		case "L":
			dx = -1
		}

		for i := 0; i < count; i++ {
			rope.Move(dx, dy)
			visited[rope.tail] = visit

			longRope.Move(dx, dy)
			longRopeVisited[longRope.body[len(longRope.body)-1]] = visit

		}
	}

	nVisited := len(visited)
	nVisitedLongRope := len(longRopeVisited)

	return Result{nVisited, nVisitedLongRope}
}
