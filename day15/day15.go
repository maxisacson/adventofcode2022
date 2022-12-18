package day15

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

func FindFreq(sensors *[]Sensor, searchMin Vec, searchMax Vec) int {
	for _, sensor := range *sensors {
		dist := sensor.dist + 1
		pos := sensor.pos
		for y := pos.y - dist; y <= pos.y+dist; y++ {
			x := dist + pos.x - Abs(y-pos.y)
			checkPos := Vec{x, y}
			if searchMin.y <= y && y < searchMax.y && searchMin.x <= x && x < searchMax.x {
				result, _, _ := IsCovered(sensors, checkPos)
				if !result {
					return 4000000*x + y
				}
			}

			x = pos.x + Abs(y-pos.y) - dist
			checkPos = Vec{x, y}
			if searchMin.y <= y && y < searchMax.y && searchMin.x <= x && x < searchMax.x {
				result, _, _ := IsCovered(sensors, checkPos)
				if !result {
					return 4000000*x + y
				}
			}
		}
	}

	return 0
}

func Run(fileName string, targetRow int, searchSize int) Result {
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
	}

	count := 0
	y := targetRow
	for x := xMin; x <= xMax; x++ {
		pos := Vec{x, y}
		result := CanContainBeacon(&sensors, pos)
		if !result {
			count++
		}
	}

	searchMin := Vec{0, 0}
	searchMax := Vec{searchSize, searchSize}
	freq := FindFreq(&sensors, searchMin, searchMax)

	return Result{count, freq}
}
