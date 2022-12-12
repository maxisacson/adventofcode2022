package day12

import (
	"aoc22/utils"
	"sort"
)

type Result struct {
	part1 int
	part2 int
}

type Vec struct {
	x int
	y int
}

func (u Vec) Add(v Vec) Vec {
	u.x += v.x
	u.y += v.y
	return u
}

func IsValid(heightMap *[][]int, pos Vec) bool {
	nRows := len(*heightMap)
	nCols := len((*heightMap)[0])

	if pos.x < 0 || pos.x >= nCols || pos.y < 0 || pos.y >= nRows {
		return false
	}

	return true
}

func FindPath(heightMap *[][]int, start, goal Vec) int {
	adjacent := []Vec{
		{0, -1},
		{0, 1},
		{1, 0},
		{-1, 0},
	}

	nRows := len(*heightMap)
	nCols := len((*heightMap)[0])
	distance := make([][]int, nRows)
	visited := make([][]int, nRows)

	unvisited := []Vec{}
	for y := range distance {
		distance[y] = make([]int, nCols)
		visited[y] = make([]int, nCols)
		for x := range distance[y] {
			distance[y][x] = 99999
			unvisited = append(unvisited, Vec{x, y})
			visited[y][x] = 0
		}
	}
	pos := start
	distance[pos.y][pos.x] = 0

	for len(unvisited) > 0 {
		sort.Slice(unvisited, func(i, j int) bool {
			u := unvisited[i]
			v := unvisited[j]
			uDist := distance[u.y][u.x]
			vDist := distance[v.y][v.x]
			return uDist > vDist
		})

		next := len(unvisited) - 1
		pos = unvisited[next]
		visited[pos.y][pos.x] = 1
		unvisited = unvisited[:next]
		height := (*heightMap)[pos.y][pos.x]

		if pos == goal {
			break
		}

		for _, v := range adjacent {
			newPos := pos.Add(v)

			if !IsValid(heightMap, newPos) {
				continue
			}

			if visited[newPos.y][newPos.x] == 1 {
				continue
			}

			newHeight := (*heightMap)[newPos.y][newPos.x]
			if newHeight > height+1 {
				continue
			}

			newDist := distance[pos.y][pos.x] + 1
			if newDist < distance[newPos.y][newPos.x] {
				distance[newPos.y][newPos.x] = newDist
			}
		}
	}

	return distance[goal.y][goal.x]
}

func Run(fileName string) Result {
	lines := utils.ReadFileToLines(fileName)
	nRows := len(lines)
	nCols := len(lines[0])

	start := Vec{}
	goal := Vec{}

	heightMap := make([][]int, nRows)
	for y, line := range lines {
		heightMap[y] = make([]int, nCols)
		for x, c := range line {
			h := int(c - 'a')
			if c == 'S' {
				start = Vec{x, y}
				h = 0
			} else if c == 'E' {
				goal = Vec{x, y}
				h = int('z' - 'a')
			}
			heightMap[y][x] = h
		}
	}

	steps := FindPath(&heightMap, start, goal)

	return Result{steps, 0}
}
