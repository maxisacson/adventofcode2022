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

func RunBlueprint(bp Blueprint) int {
	ore := 0
	clay := 0
	obsidian := 0
	geodes := 0

	nOreRobots := 1
	nClayRobots := 0
	nObsidianRobots := 0
	nGeodeRobots := 0

	nOreRobotsQueued := 0
	nClayRobotsQueued := 0
	nObsidianRobotsQueued := 0
	nGeodeRobotsQueued := 0

	for i := 1; i <= 24; i++ {
		if ore >= bp.GeodeRobotCost.first && obsidian >= bp.GeodeRobotCost.second {
			ore -= bp.GeodeRobotCost.first
			obsidian -= bp.GeodeRobotCost.second
			nGeodeRobotsQueued++
		}

		if ore >= bp.ObsidianRobotCost.first && clay >= bp.ObsidianRobotCost.second {
			ore -= bp.ObsidianRobotCost.first
			clay -= bp.ObsidianRobotCost.second
			nObsidianRobotsQueued++
		}

		if ore >= bp.ClayRobotCost {
			ore -= bp.ClayRobotCost
			nClayRobotsQueued++
		}

		if ore >= bp.OreRobotCost {
			ore -= bp.OreRobotCost
			nOreRobotsQueued++
		}

		ore += nOreRobots
		clay += nClayRobots
		obsidian += nObsidianRobots
		geodes += nGeodeRobots

		nGeodeRobots += nGeodeRobotsQueued
		nObsidianRobots += nObsidianRobotsQueued
		nClayRobots += nClayRobotsQueued
		nOreRobots += nOreRobotsQueued

		nGeodeRobotsQueued = 0
		nObsidianRobotsQueued = 0
		nClayRobotsQueued = 0
		nOreRobotsQueued = 0

		fmt.Printf("ore:%d clay:%d obs:%d geo:%d  ore rob:%d clay rob:%d obs rob:%d geo rob:%d\n", ore, clay, obsidian, geodes, nOreRobots, nClayRobots, nObsidianRobots, nGeodeRobots)
	}

	return geodes
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

	for _, bp := range blueprints {
		count := RunBlueprint(bp)
		fmt.Println(count)
	}

	return Result{len(lines), len(lines)}
}
