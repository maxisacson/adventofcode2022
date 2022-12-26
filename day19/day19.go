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

func (bp Blueprint) String() string {
	return fmt.Sprintf(
		"OreRobotCost: %v\n"+
			"ClayRobotCost: %v\n"+
			"ObsidianRobotCost: %v\n"+
			"GeodeRobotCost: %v", bp.OreRobotCost, bp.ClayRobotCost, bp.ObsidianRobotCost, bp.GeodeRobotCost)
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

func (s State) String() string {
	str := ""

	str += fmt.Sprintf("ore:      %d\n", s.ore)
	str += fmt.Sprintf("clay:     %d\n", s.clay)
	str += fmt.Sprintf("obsidian: %d\n", s.obsidian)
	str += fmt.Sprintf("geodes:   %d\n", s.geodes)

	str += "\n"

	str += fmt.Sprintf("nOreRobots:      %d\n", s.nOreRobots)
	str += fmt.Sprintf("nClayRobots:     %d\n", s.nClayRobots)
	str += fmt.Sprintf("nObsidianRobots: %d\n", s.nObsidianRobots)
	str += fmt.Sprintf("nGeodeRobots:    %d\n", s.nGeodeRobots)

	return str
}

func Collect(state *State) {
	state.ore += state.nOreRobots
	state.clay += state.nClayRobots
	state.obsidian += state.nObsidianRobots
	state.geodes += state.nGeodeRobots

	state.nOreRobots += state.nOreRobotsQueued
	state.nClayRobots += state.nClayRobotsQueued
	state.nObsidianRobots += state.nObsidianRobotsQueued
	state.nGeodeRobots += state.nGeodeRobotsQueued

	state.nOreRobotsQueued = 0
	state.nClayRobotsQueued = 0
	state.nObsidianRobotsQueued = 0
	state.nGeodeRobotsQueued = 0
}

func Max(a, b int) int {
	if b > a {
		return b
	}

	return a
}

// func RunBlueprint(timeLeft int, bp Blueprint, state State) int {
// 	if timeLeft == 0 {
// 		return 0
// 	}
//
// 	// fmt.Println(timeLeft)
// 	// fmt.Println(state)
//
// 	maxGeodes := 0
//
// 	if state.ore >= bp.GeodeRobotCost.first && state.obsidian >= bp.GeodeRobotCost.second {
// 		state.nGeodeRobotsQueued++
// 		state.ore -= bp.GeodeRobotCost.first
// 		state.obsidian -= bp.GeodeRobotCost.second
// 	}
//
// 	if state.ore >= bp.ObsidianRobotCost.first && state.clay >= bp.ObsidianRobotCost.second {
// 		state.nObsidianRobotsQueued++
// 		state.ore -= bp.ObsidianRobotCost.first
// 		state.clay -= bp.ObsidianRobotCost.second
// 	}
//
// 	if state.ore >= bp.OreRobotCost {
// 		next := state
// 		next.nOreRobotsQueued++
// 		next.ore -= bp.OreRobotCost
// 		geodes := Round(&next)
// 		geodes += RunBlueprint(timeLeft-1, bp, next)
// 		maxGeodes = Max(maxGeodes, geodes)
// 	}
//
// 	if state.ore >= bp.ClayRobotCost {
// 		next := state
// 		next.nClayRobotsQueued++
// 		next.ore -= bp.ClayRobotCost
// 		geodes := Round(&next)
// 		geodes += RunBlueprint(timeLeft-1, bp, next)
// 		maxGeodes = Max(maxGeodes, geodes)
// 	}
//
// 	geodes := Round(&state)
// 	geodes += RunBlueprint(timeLeft-1, bp, state)
// 	maxGeodes = Max(maxGeodes, geodes)
//
// 	return maxGeodes
// }

func QueueGeodeRobot(state *State, bp Blueprint) {
	state.nGeodeRobotsQueued++
	state.ore -= bp.GeodeRobotCost.first
	state.obsidian -= bp.GeodeRobotCost.second
	// fmt.Println("queue geode robot")
}

func QueueObsidianRobot(state *State, bp Blueprint) {
	state.nObsidianRobotsQueued++
	state.ore -= bp.ObsidianRobotCost.first
	state.clay -= bp.ObsidianRobotCost.second
	// fmt.Println("queue obsidian robot")
}

func QueueClayRobot(state *State, bp Blueprint) {
	state.nClayRobotsQueued++
	state.ore -= bp.ClayRobotCost
	// fmt.Println("queue clay robot")
}

func QueueOreRobot(state *State, bp Blueprint) {
	state.nOreRobotsQueued++
	state.ore -= bp.OreRobotCost
	// fmt.Println("queue ore robot")
}

func CanBuildGeodeRobot(state *State, bp Blueprint) bool {
	return state.ore >= bp.GeodeRobotCost.first && state.obsidian >= bp.GeodeRobotCost.second
}

func CanBuildObsidianRobot(state *State, bp Blueprint) bool {
	return state.ore >= bp.ObsidianRobotCost.first && state.clay >= bp.ObsidianRobotCost.second
}

func CanBuildClayRobot(state *State, bp Blueprint) bool {
	return state.ore >= bp.ClayRobotCost
}

func CanBuildOreRobot(state *State, bp Blueprint) bool {
	return state.ore >= bp.OreRobotCost
}

// func RunBlueprint(timeLeft int, bp Blueprint, state *State) {
// 	if timeLeft == 0 {
// 		return
// 	}
//
// 	// fmt.Println(state)
// 	fmt.Println("=====", 24-timeLeft+1, "=====")
// 	fmt.Println(state)
//
// 	if state.ore >= bp.GeodeRobotCost.first && state.obsidian >= bp.GeodeRobotCost.second {
// 		QueueGeodeRobot(state, bp)
// 	}
//
// 	if state.ore >= bp.ObsidianRobotCost.first && state.clay >= bp.ObsidianRobotCost.second {
// 		if state.nObsidianRobots > 0 {
// 			nRoundsObsidian := (bp.GeodeRobotCost.second - state.obsidian) / state.nObsidianRobots
// 			nRoundsOre := (bp.GeodeRobotCost.first - state.ore + bp.ObsidianRobotCost.first) / state.nOreRobots
// 			if nRoundsOre <= nRoundsObsidian {
// 				QueueObsidianRobot(state, bp)
// 			}
// 		} else {
// 			QueueObsidianRobot(state, bp)
// 		}
// 	}
//
// 	if state.ore >= bp.ClayRobotCost {
// 		if state.nClayRobots > 0 {
// 			nRoundsClay := (bp.ObsidianRobotCost.second - state.clay) / state.nClayRobots
// 			nRoundsOre := (bp.ObsidianRobotCost.first - state.ore + bp.ClayRobotCost) / state.nOreRobots
// 			if nRoundsOre <= nRoundsClay {
// 				QueueClayRobot(state, bp)
// 			}
// 		} else {
// 			QueueClayRobot(state, bp)
// 		}
// 	}
//
// 	if state.ore >= bp.OreRobotCost {
// 		nRoundsOre := (bp.ClayRobotCost - state.ore) / state.nOreRobots
// 		nRoundsOre2 := (bp.ClayRobotCost - state.ore + bp.OreRobotCost) / state.nOreRobots
// 		if nRoundsOre2 <= nRoundsOre {
// 			QueueOreRobot(state, bp)
// 		}
// 	}
//
// 	Collect(state)
// 	RunBlueprint(timeLeft-1, bp, state)
// }

func RunBlueprint(timeLeft int, bp Blueprint, state *State) {
	if timeLeft == 0 {
		return
	}

	current := *state
	Collect(state)
	RunBlueprint(timeLeft-1, bp, state)

	if CanBuildGeodeRobot(&current, bp) {
		next := current
		QueueGeodeRobot(&next, bp)
		Collect(&next)
		RunBlueprint(timeLeft-1, bp, &next)
		if next.geodes > state.geodes {
			*state = next
		}
	}

	if CanBuildObsidianRobot(&current, bp) {
		next := current
		QueueObsidianRobot(&next, bp)
		Collect(&next)
		RunBlueprint(timeLeft-1, bp, &next)
		if next.geodes > state.geodes {
			*state = next
		}
	}

	if CanBuildClayRobot(&current, bp) {
		next := current
		QueueClayRobot(&next, bp)
		Collect(&next)
		RunBlueprint(timeLeft-1, bp, &next)
		if next.geodes > state.geodes {
			*state = next
		}
	}

	// if CanBuildOreRobot(&current, bp) && current.nClayRobots > 0 && current.nObsidianRobots > 0 && current.nGeodeRobots > 0 { //&& current.nOreRobots < 4 { // && current.nOreRobots < current.nClayRobots { //!CanBuildClayRobot(&current, bp) { // && !CanBuildGeodeRobot(&current, bp) && !CanBuildObsidianRobot(&current, bp) {
	if CanBuildOreRobot(&current, bp) {
		next := current
		QueueOreRobot(&next, bp)
		Collect(&next)
		RunBlueprint(timeLeft-1, bp, &next)
		if next.geodes > state.geodes {
			*state = next
		}
	}
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
	// RunBlueprint(24, blueprints[0], &state)
	// fmt.Println(state)

	sum := 0
	for i, bp := range blueprints {
		thisState := state
		fmt.Printf("blueprint %d/%d\n", i+1, len(blueprints))
		fmt.Println(bp)
		RunBlueprint(24, bp, &thisState)
		fmt.Println(thisState)
		sum += (i + 1) * thisState.geodes
		// break
	}

	return Result{sum, 0}
}
