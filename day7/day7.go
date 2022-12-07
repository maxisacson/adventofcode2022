package day7

import (
	"aoc22/utils"
	"fmt"
	"strconv"
	"strings"
)

func Run(fileName string) int {
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
					}
					cwd = string(line[5:])
					if cwd == ".." {
						cwd = parent
						stack = stack[:top]
					} else {
						if cwd == "/" {
							stack = make([]string, 0)
						}
						stack = append(stack, parent+"/"+cwd)
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
	for _, val := range sizes {
		if val > 100000 {
			continue
		}
		sum += val
	}

	return sum
}
