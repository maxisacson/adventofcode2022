package dayN

import (
	"aoc22/utils"
)

type Result struct {
	part1 int
	part2 int
}

func Run(fileName string) Result {
	lines := utils.ReadFileToLines(fileName)

	return Result{len(lines), len(lines)}
}
