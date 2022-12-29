package day16

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

type Valve struct {
	name     string
	flowRate int
	tunnels  []string
	isOpen   bool
	openTime int
}

type Cave struct {
	pos      string
	valves   map[string]Valve
	timeLeft int
	pressure int
}

func AllValvesOpen(valves map[string]Valve) bool {
	for _, v := range valves {
		if v.flowRate > 0 && !v.isOpen {
			return false
		}
	}

	return true
}

func copyValves(valves map[string]Valve) map[string]Valve {
	ret := make(map[string]Valve)
	for k, v := range valves {
		ret[k] = v
	}
	return ret
}

func copyInto(dst *map[string]Valve, src map[string]Valve) {
	for k, v := range src {
		(*dst)[k] = v
	}
}

func Traverse(valves map[string]Valve, name string, timeLeft int, path *[]string) int {
	if timeLeft <= 0 {
		return 0
	}

	if AllValvesOpen(valves) {
		return 0
	}

	valve := valves[name]
	valvesOpen := copyValves(valves)
	valvesClosed := copyValves(valves)
	*path = append(*path, fmt.Sprintf("%s:%d", name, timeLeft))
	newPath := []string{}

	maxPressure := -1
	if !valve.isOpen && valve.flowRate > 0 {
		pressure := valve.flowRate * (timeLeft - 1)
		valve.isOpen = true
		valvesOpen[name] = valve
		valvesCopy := copyValves(valvesOpen)

		pathCopy := make([]string, len(*path))
		copy(pathCopy, *path)
		pathCopy = append(pathCopy, fmt.Sprintf("%s:%d", name, timeLeft-1))
		// fmt.Println(pathCopy)

		for _, next := range valve.tunnels {
			thisPressure := pressure + Traverse(valvesCopy, next, timeLeft-2, &pathCopy)
			if thisPressure > maxPressure {
				maxPressure = thisPressure
				copyInto(&valves, valvesCopy)
				newPath = pathCopy
			}
		}
	}

	for _, next := range valve.tunnels {
		valvesCopy := copyValves(valvesClosed)

		pathCopy := make([]string, len(*path))
		copy(pathCopy, *path)
		// fmt.Println(pathCopy)

		thisPressure := Traverse(valvesCopy, next, timeLeft-1, &pathCopy)
		if thisPressure > maxPressure {
			maxPressure = thisPressure
			copyInto(&valves, valvesCopy)
			newPath = pathCopy
		}
	}

	// for _, v := range valves {
	// 	fmt.Println(v.name, v.isOpen)
	// }
	// fmt.Println()
	*path = newPath
	return maxPressure
}

func Run(fileName string) Result {
	lines := utils.ReadFileToLines(fileName)

	valves := make(map[string]Valve)

	for _, line := range lines {
		parts := strings.Split(line, "; ")
		fields0 := strings.Fields(parts[0])
		valve := fields0[1]
		flowRate, _ := strconv.Atoi(strings.Split(fields0[4], "=")[1])
		connections := strings.Split(parts[1][23:len(parts[1])], ", ")
		valves[valve] = Valve{valve, flowRate, connections, false, 0}
		fmt.Println(valves[valve])
	}
	fmt.Println()
	cave := Cave{}
	cave.valves = valves
	cave.pos = "AA"
	cave.timeLeft = 24

	pressure := 0
	path := []string{}
	pressure = Traverse(valves, "AA", 5, &path)

	fmt.Println(path)
	for _, v := range valves {
		fmt.Println(v.name, v.isOpen)
	}
	fmt.Println()

	return Result{pressure, 0}
}
