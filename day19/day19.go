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

type Cost struct {
	Ore      int
	Clay     int
	Obsidian int
}

type Blueprint struct {
	OreRobotCost      Cost
	ClayRobotCost     Cost
	ObsidianRobotCost Cost
	GeodeRobotCost    Cost
}

func (bp Blueprint) String() string {
	return fmt.Sprintf(
		"OreRobotCost: %d\n"+
			"ClayRobotCost: %d\n"+
			"ObsidianRobotCost: %d %d\n"+
			"GeodeRobotCost: %d %d",
		bp.OreRobotCost.Ore,
		bp.ClayRobotCost.Ore,
		bp.ObsidianRobotCost.Ore,
		bp.ObsidianRobotCost.Clay,
		bp.GeodeRobotCost.Ore,
		bp.GeodeRobotCost.Obsidian)
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

func QueueGeodeRobot(state *State, bp Blueprint) {
	state.nGeodeRobotsQueued++
	state.ore -= bp.GeodeRobotCost.Ore
	state.obsidian -= bp.GeodeRobotCost.Obsidian
	// fmt.Println("queue geode robot")
}

func QueueObsidianRobot(state *State, bp Blueprint) {
	state.nObsidianRobotsQueued++
	state.ore -= bp.ObsidianRobotCost.Ore
	state.clay -= bp.ObsidianRobotCost.Clay
	// fmt.Println("queue obsidian robot")
}

func QueueClayRobot(state *State, bp Blueprint) {
	state.nClayRobotsQueued++
	state.ore -= bp.ClayRobotCost.Ore
	// fmt.Println("queue clay robot")
}

func QueueOreRobot(state *State, bp Blueprint) {
	state.nOreRobotsQueued++
	state.ore -= bp.OreRobotCost.Ore
	// fmt.Println("queue ore robot")
}

func CanBuildGeodeRobot(state State, bp Blueprint) bool {
	return state.ore >= bp.GeodeRobotCost.Ore && state.obsidian >= bp.GeodeRobotCost.Obsidian
}

func CanBuildObsidianRobot(state State, bp Blueprint) bool {
	return state.ore >= bp.ObsidianRobotCost.Ore && state.clay >= bp.ObsidianRobotCost.Clay
}

func CanBuildClayRobot(state State, bp Blueprint) bool {
	return state.ore >= bp.ClayRobotCost.Ore
}

func CanBuildOreRobot(state State, bp Blueprint) bool {
	return state.ore >= bp.OreRobotCost.Ore
}

func CanBuildGeodeRobotNextRound(state State, bp Blueprint) bool {
	return bp.GeodeRobotCost.Obsidian <= state.obsidian+state.nObsidianRobots && bp.GeodeRobotCost.Ore <= state.ore+state.nOreRobots
}

func RoundsUntilNextGeodeRobot(state State, bp Blueprint) int {
	if state.nOreRobots == 0 || state.nObsidianRobots == 0 {
		return 9999
	}

	ore := state.ore
	obsidian := state.obsidian

	rounds := 0
	for ore < bp.GeodeRobotCost.Ore && obsidian < bp.GeodeRobotCost.Obsidian {
		rounds++
		ore += state.nOreRobots
		obsidian += state.nObsidianRobots
	}
	return rounds
}

func RunBlueprint(timeLeft int, bp Blueprint, state *State) {
	if timeLeft == 0 {
		return
	}

	if CanBuildGeodeRobot(*state, bp) {
		QueueGeodeRobot(state, bp)
		Collect(state)
		RunBlueprint(timeLeft-1, bp, state)
	} else {
		current := *state
		Collect(state)
		RunBlueprint(timeLeft-1, bp, state)

		if CanBuildObsidianRobot(current, bp) { // && !CanBuildGeodeRobotNextRound(current, bp) {
			next := current
			QueueObsidianRobot(&next, bp)
			Collect(&next)
			RunBlueprint(timeLeft-1, bp, &next)
			if next.geodes > state.geodes {
				*state = next
			}
		}

		if CanBuildClayRobot(current, bp) { //&& !CanBuildGeodeRobotNextRound(current, bp) {
			next := current
			QueueClayRobot(&next, bp)
			Collect(&next)
			RunBlueprint(timeLeft-1, bp, &next)
			if next.geodes > state.geodes {
				*state = next
			}
		}

		if CanBuildOreRobot(current, bp) { //&& !CanBuildGeodeRobotNextRound(current, bp) {
			next := current
			QueueOreRobot(&next, bp)
			Collect(&next)
			RunBlueprint(timeLeft-1, bp, &next)
			if next.geodes > state.geodes {
				*state = next
			}
		}
	}
}

func Run(fileName string) Result {
	lines := utils.ReadFileToLines(fileName)

	blueprints := []Blueprint{}
	for _, line := range lines {
		parts := strings.Split(line, ":")
		costStrings := strings.Split(parts[1], ".")
		oreRobotCostOre, _ := strconv.Atoi(strings.Fields(costStrings[0])[4])
		clayRobotCostOre, _ := strconv.Atoi(strings.Fields(costStrings[1])[4])
		obsidianRobotCostOre, _ := strconv.Atoi(strings.Fields(costStrings[2])[4])
		obsidianRobotCostClay, _ := strconv.Atoi(strings.Fields(costStrings[2])[7])
		geodeRobotCostOre, _ := strconv.Atoi(strings.Fields(costStrings[3])[4])
		geodeRobotCostObsidian, _ := strconv.Atoi(strings.Fields(costStrings[3])[7])

		blueprint := Blueprint{
			Cost{oreRobotCostOre, 0, 0},
			Cost{clayRobotCostOre, 0, 0},
			Cost{obsidianRobotCostOre, obsidianRobotCostClay, 0},
			Cost{geodeRobotCostOre, 0, geodeRobotCostObsidian},
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
		break
	}

	return Result{sum, 0}
}
