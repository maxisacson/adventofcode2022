package day15

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

type Sensor struct {
	pos    Vec
	beacon Vec
	dist   int
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Dist(u Vec, v Vec) int {
	return Abs(u.x-v.x) + Abs(u.y-v.y)
}

func Min(a int, b int) int {
	if b < a {
		return b
	}
	return a
}

func Max(a int, b int) int {
	if b > a {
		return b
	}
	return a
}

func ParseSensor(line string) Sensor {
	fields := strings.Fields(line)
	xs, _ := strconv.Atoi(fields[2][2 : len(fields[2])-1])
	ys, _ := strconv.Atoi(fields[3][2 : len(fields[3])-1])
	xb, _ := strconv.Atoi(fields[8][2 : len(fields[8])-1])
	yb, _ := strconv.Atoi(fields[9][2:len(fields[9])])

	pos := Vec{xs, ys}
	beacon := Vec{xb, yb}

	return Sensor{pos, beacon, Dist(pos, beacon)}
}

func Print() {
	fmt.Println()
}

func FindClosest(sensors *[]Sensor, pos Vec) (Sensor, int) {
	minIndex := 0
	minDist := 9999
	for i, sensor := range *sensors {
		dist := Dist(sensor.pos, pos)
		if dist < minDist {
			minDist = dist
			minIndex = i
		}
	}

	return (*sensors)[minIndex], minDist
}

func IsCovered(sensors *[]Sensor, pos Vec) (bool, Sensor, int) {
	minIndex := 0
	minDist := 9999

	result := false
	for i, sensor := range *sensors {
		dist := Dist(sensor.pos, pos)

		if dist <= sensor.dist {
			result = true
		}

		if dist < minDist {
			minDist = dist
			minIndex = i
		}
	}

	return result, (*sensors)[minIndex], minDist
}

func CanContainBeacon(sensors *[]Sensor, pos Vec) bool {
	result := true
	for _, sensor := range *sensors {
		dist := Dist(sensor.pos, pos)
		if dist <= sensor.dist && pos != sensor.beacon {
			return false
		}
	}

	return result
}

func Run(fileName string, targetRow int) Result {
	lines := utils.ReadFileToLines(fileName)

	xMin := 9999
	xMax := -xMin
	yMin := 9999
	yMax := -yMin

	sensors := []Sensor{}
	for _, line := range lines {
		sensor := ParseSensor(line)
		sensors = append(sensors, sensor)
		xMin = Min(xMin, Min(sensor.pos.x-sensor.dist, sensor.beacon.x-sensor.dist))
		xMax = Max(xMax, Max(sensor.pos.x+sensor.dist, sensor.beacon.x+sensor.dist))
		yMin = Min(yMin, Min(sensor.pos.y-sensor.dist, sensor.beacon.y-sensor.dist))
		yMax = Max(yMax, Max(sensor.pos.y+sensor.dist, sensor.beacon.y+sensor.dist))
		// fmt.Println(sensor)
	}

	// nRows := yMax - yMin + 1
	// nCols := xMax - xMin + 1

	// coverage := make([][]bool, nRows)
	// for i := range coverage {
	// 	coverage[i] = make([]bool, nCols)
	// }
	//
	// for _, sensor := range sensors {
	// 	x0 := sensor.pos.x - sensor.dist
	// 	y0 := sensor.pos.y - sensor.dist
	// 	x1 := sensor.pos.x + sensor.dist
	// 	y1 := sensor.pos.y + sensor.dist
	// 	for y := y0; y <= y1; y++ {
	// 		for x := x0; x <= x1; x++ {
	// 			i := y - yMin
	// 			j := x - xMin
	// 			pos := Vec{x, y}
	// 			dist := Dist(sensor.pos, pos)
	// 			if dist <= sensor.dist && pos != sensor.beacon {
	// 				coverage[i][j] = true
	// 			}
	// 		}
	// 	}
	// }

	// sensors2 := []Sensor{sensors[6]}
	//
	// for y := yMin; y <= yMax; y++ {
	// 	for x := xMin; x <= xMax; x++ {
	// 		pos := Vec{x, y}
	// 		result, sensor, _ := IsCovered(&sensors2, pos)
	// 		if !result {
	// 			fmt.Print(".")
	// 		} else if pos == sensor.pos {
	// 			fmt.Print("S")
	// 		} else if pos == sensor.beacon {
	// 			fmt.Print("B")
	// 		} else {
	// 			fmt.Print("#")
	// 		}
	// 	}
	// 	fmt.Println()
	// }

	count := 0
	y := targetRow
	for x := xMin; x <= xMax; x++ {
		pos := Vec{x, y}
		result := CanContainBeacon(&sensors, pos)
		if !result {
			count++
		}
	}

	return Result{count, 0}
}
