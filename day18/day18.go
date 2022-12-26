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

	min Vec
	max Vec
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func Max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func (c *Cluster) AddCube(pos Vec) {
	count := c.CountAdjecent(pos)
	c.area = c.area + 6 - 2*count
	c.cubes = append(c.cubes, pos)

	c.min.x = Min(c.min.x, pos.x)
	c.min.y = Min(c.min.y, pos.y)
	c.min.z = Min(c.min.z, pos.z)

	c.max.x = Max(c.max.x, pos.x)
	c.max.y = Max(c.max.y, pos.y)
	c.max.z = Max(c.max.z, pos.z)
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

func (c *Cluster) Exterior() int {
	nz := c.max.z - c.min.z + 1
	ny := c.max.y - c.min.y + 1
	nx := c.max.x - c.min.x + 1

	cubeMap := make([][][]int, nz)
	for i := range cubeMap {
		cubeMap[i] = make([][]int, ny)
		for j := range cubeMap[i] {
			cubeMap[i][j] = make([]int, nx)
			for k := range cubeMap[i][j] {
				cubeMap[i][j][k] = -1
			}
		}
	}

	for ind, cube := range c.cubes {
		i := cube.z - c.min.z
		j := cube.y - c.min.y
		k := cube.x - c.min.x
		cubeMap[i][j][k] = ind
	}

	done := false
	for !done {
		done = true
		for i := range cubeMap {
			for j := range cubeMap[i] {
				for k := range cubeMap[i][j] {
					ind := cubeMap[i][j][k]
					if ind >= 0 || ind == -2 {
						continue
					}

					if i == 0 || i == nz-1 || j == 0 || j == ny-1 || k == 0 || k == nx-1 ||
						cubeMap[i-1][j][k] == -2 || cubeMap[i+1][j][k] == -2 ||
						cubeMap[i][j-1][k] == -2 || cubeMap[i][j+1][k] == -2 ||
						cubeMap[i][j][k-1] == -2 || cubeMap[i][j][k+1] == -2 {
						// exterior point
						cubeMap[i][j][k] = -2
						done = false
						continue
					}
				}
			}
		}
	}

	tmpCluster := Cluster{}

	for i := range cubeMap {
		for j := range cubeMap[i] {
			for k := range cubeMap[i][j] {
				ind := cubeMap[i][j][k]
				if ind >= 0 {
					tmpCluster.AddCube(c.cubes[ind])
				} else if ind == -1 {
					cube := Vec{}
					cube.x = k + c.min.x
					cube.y = j + c.min.y
					cube.z = i + c.min.z
					tmpCluster.AddCube(cube)
				}
			}
		}
	}

	return tmpCluster.area
}

func Run(fileName string) Result {
	lines := utils.ReadFileToLines(fileName)

	cubes := []Vec{}
	for _, line := range lines {
		xyz := strings.Split(line, ",")
		x, _ := strconv.Atoi(xyz[0])
		y, _ := strconv.Atoi(xyz[1])
		z, _ := strconv.Atoi(xyz[2])
		cubes = append(cubes, Vec{x, y, z})
	}

	cluster := Cluster{}
	cluster.min = cubes[0]
	cluster.max = cubes[0]

	for _, cube := range cubes {
		cluster.AddCube(cube)
	}

	ext := cluster.Exterior()

	return Result{cluster.area, ext}
}
