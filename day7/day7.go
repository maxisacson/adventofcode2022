package day7

import (
	"aoc22/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Result struct {
	part1 int
	part2 int
}

type Dir struct {
	name string
	size int
}

func Run(fileName string) Result {
	lines := utils.ReadFileToLines(fileName)

	stack := make([]string, 0)
	cwd := ""
	sizes := make(map[string]int)

	for _, line := range lines {
		if line[0] == '$' {
			switch {
			case line[2:4] == "cd":
				{
					top := len(stack) - 1
					parent := ""
					if top >= 0 {
						parent = stack[top]
						if parent != "/" {
							parent += "/"
						}
					}
					cwd = string(line[5:])
					if cwd == ".." {
						cwd = parent
						stack = stack[:top]
					} else {
						if cwd == "/" {
							parent = ""
							stack = make([]string, 0)
						}
						stack = append(stack, parent+cwd)
					}
				}
				break
			case line[2:4] == "ls":
				break
			case line[:3] == "dir":
				break
			default:
				panic(fmt.Sprintf("expected command: %s\n", line))
			}
		} else if line[:3] == "dir" {

		} else {
			fields := strings.Fields(line)
			size, e := strconv.Atoi(fields[0])
			if e != nil {
				panic(e)
			}

			for i := len(stack) - 1; i >= 0; i-- {
				sizes[stack[i]] += size
			}
		}
	}

	sum := 0
	dirs := make([]Dir, 0)
	for key, val := range sizes {
		dirs = append(dirs, Dir{name: key, size: val})
		if val > 100000 {
			continue
		}
		sum += val
	}

	sort.Slice(dirs, func(i, j int) bool {
		return dirs[i].size < dirs[j].size
	})

	totalSize := 70000000
	freeSpace := totalSize - sizes["/"]
	minSpace := 30000000
	size := 0
	for _, d := range dirs {
		if freeSpace+d.size < minSpace {
			continue
		}
		size = d.size
		break
	}

	return Result{sum, size}
}
