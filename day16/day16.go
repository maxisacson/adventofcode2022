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

func AllValvesOpen(valves *map[string]Valve) bool {
	for _, v := range *valves {
		if v.flowRate > 0 && !v.isOpen {
			return false
		}
	}

	return true
}

func FindPath(valves *map[string]Valve, name string, timeLeft int, targetPressure int) int {

}

func Traverse(valves *map[string]Valve, name string, timeLeft int) int {
	pressure := 0

	if timeLeft <= 0 {
		return pressure
	}

	if AllValvesOpen(valves) {
		return pressure
	}

	valve := (*valves)[name]

	// Potential pressure release for this valve
	localPressure := valve.flowRate * (timeLeft - 1)

	maxPressure := -1
	for _, next := range valve.tunnels {
		thisPressure := FindPath(valves, next, timeLeft-1, localPressure)
		if thisPressure > maxPressure {
			maxPressure = thisPressure
		}
	}

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

	pressure := Traverse(&valves, "AA", 13)

	for _, v := range valves {
		fmt.Println(v)
	}
	fmt.Println()

	return Result{pressure, 0}
}
