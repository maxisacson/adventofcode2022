package day19

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

type Pair struct {
	first  int
	second int
}

type Blueprint struct {
	OreRobotCost      int
	ClayRobotCost     int
	ObsidianRobotCost Pair
	GeodeRobotCost    Pair
}

type State struct {
	ore      int
	clay     int
	obsidian int
	geodes   int

	nOreRobots      int
	nClayRobots     int
	nObsidianRobots int
	nGeodeRobots    int

	nOreRobotsQueued      int
	nClayRobotsQueued     int
	nObsidianRobotsQueued int
	nGeodeRobotsQueued    int
}

func RunBlueprint(timeLeft int, bp Blueprint, state State, geodesTotal int) int {
	if timeLeft == 0 {
		return 0
	}

	// queue robots
	// collect resources
	// build robot

	states := []State{}

	maxGeodes := 0
	if state.ore >= bp.OreRobotCost {
		nextState := state
		nextState.ore -= bp.OreRobotCost
		nextState.nOreRobotsQueued++
		states = append(states, nextState)
	}

	if state.clay >= bp.ClayRobotCost {
		nextState := state
		nextState.clay -= bp.ClayRobotCost
		nextState.nClayRobotsQueued++
		states = append(states, nextState)
	}

	if state.ore >= bp.ObsidianRobotCost.first && state.clay >= bp.ObsidianRobotCost.second {
		nextState := state
		nextState.ore -= bp.ObsidianRobotCost.first
		nextState.clay -= bp.ObsidianRobotCost.second
		nextState.nObsidianRobotsQueued++
		states = append(states, nextState)
	}

	if state.ore >= bp.GeodeRobotCost.first && state.obsidian >= bp.GeodeRobotCost.second {
		nextState := state
		nextState.ore -= bp.GeodeRobotCost.first
		nextState.obsidian -= bp.GeodeRobotCost.second
		nextState.nGeodeRobotsQueued++
		states = append(states, nextState)
	}

	for _, nextState := range states {
		nextState.ore += nextState.nOreRobots
		nextState.clay += nextState.nClayRobots
		nextState.obsidian += nextState.nObsidianRobots
		nextState.geodes += nextState.nGeodeRobots

		nextState.nOreRobots += nextState.nOreRobotsQueued
		nextState.nClayRobots += nextState.nClayRobotsQueued
		nextState.nObsidianRobots += nextState.nObsidianRobotsQueued
		nextState.nGeodeRobots += nextState.nGeodeRobotsQueued

		nextState.nOreRobotsQueued = 0
		nextState.nClayRobotsQueued = 0
		nextState.nObsidianRobotsQueued = 0
		nextState.nGeodeRobotsQueued = 0

		geodes := RunBlueprint(timeLeft-1, bp, nextState, nextState.geodes)
		if geodes > maxGeodes {
			maxGeodes = geodes
		}
	}

	return maxGeodes + geodesTotal
}

func Run(fileName string) Result {
	lines := utils.ReadFileToLines(fileName)

	blueprints := []Blueprint{}
	for _, line := range lines {
		parts := strings.Split(line, ":")
		costStrings := strings.Split(parts[1], ".")
		oreRobotCost, _ := strconv.Atoi(strings.Fields(costStrings[0])[4])
		clayRobotCost, _ := strconv.Atoi(strings.Fields(costStrings[1])[4])
		obsidianRobotCostFirst, _ := strconv.Atoi(strings.Fields(costStrings[2])[4])
		obsidianRobotCostSecond, _ := strconv.Atoi(strings.Fields(costStrings[2])[7])
		geodeRobotCostFirst, _ := strconv.Atoi(strings.Fields(costStrings[3])[4])
		geodeRobotCostSecond, _ := strconv.Atoi(strings.Fields(costStrings[3])[7])

		blueprint := Blueprint{
			oreRobotCost,
			clayRobotCost,
			Pair{obsidianRobotCostFirst, obsidianRobotCostSecond},
			Pair{geodeRobotCostFirst, geodeRobotCostSecond},
		}
		blueprints = append(blueprints, blueprint)
	}

	state := State{}

	state.nOreRobots = 1

	for _, bp := range blueprints {
		count := RunBlueprint(24, bp, state, 0)
		fmt.Println(count)
	}

	return Result{len(lines), len(lines)}
}
