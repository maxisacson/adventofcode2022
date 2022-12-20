package day18

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
	z int
}

type Cluster struct {
	cubes []Vec
	area  int
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (c *Cluster) AddCube(pos Vec) {
	count := c.CountAdjecent(pos)
	c.area = c.area + 6 - 2*count
	c.cubes = append(c.cubes, pos)
}

func IsAdjecent(pos1, pos2 Vec) bool {
	dist := Abs(pos2.x-pos1.x) + Abs(pos2.y-pos1.y) + Abs(pos2.z-pos1.z)
	return (dist == 1)
}

func (c *Cluster) CountAdjecent(pos Vec) int {
	count := 0

	for _, cube := range c.cubes {
		if IsAdjecent(pos, cube) {
			count++
		}
	}

	return count
}

func Run(fileName string) Result {
	lines := utils.ReadFileToLines(fileName)

	cluster := Cluster{}
	for _, line := range lines {
		xyz := strings.Split(line, ",")
		x, _ := strconv.Atoi(xyz[0])
		y, _ := strconv.Atoi(xyz[1])
		z, _ := strconv.Atoi(xyz[2])
		cluster.AddCube(Vec{x, y, z})
	}

	return Result{cluster.area, 0}
}
