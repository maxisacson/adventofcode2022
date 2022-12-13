package day13

import (
	"aoc22/utils"
	"bytes"
	"fmt"
	"strconv"
)

type Result struct {
	part1 int
	part2 int
}

const (
	List int = iota
	Item     = iota
)

type ListItem struct {
	Type  int
	List  []ListItem
	Value int
}

func (item ListItem) String() string {
	if item.Type == Item {
		return fmt.Sprint(item.Value)
	}

	return fmt.Sprint(item.List)
}

func ParseList(line string) []ListItem {
	list := []ListItem{}

	i := 0
	for i < len(line) {
		c := line[i]
		if c == '[' && i > 0 {
			// Find sublist
			j := i
			count := 1
			for count > 0 {
				j++
				c = line[j]
				if c == ']' {
					count--
				} else if c == '[' {
					count++
				}
			}
			item := ListItem{}
			item.Type = List
			item.List = ParseList(line[i : j+1])
			list = append(list, item)
			i = j + 1
		} else if c != '[' && c != ']' {
			// Parse int
			var bytes bytes.Buffer
			for c >= '0' && c <= '9' {
				bytes.WriteByte(c)
				i++
				c = line[i]
			}
			item := ListItem{}
			item.Type = Item
			item.Value, _ = strconv.Atoi(bytes.String())
			list = append(list, item)
		}
		i++
	}

	return list
}

func Compare(a, b ListItem) int {
	if a.Type == Item && b.Type == Item {
		if a.Value < b.Value {
			return 1
		} else if b.Value < a.Value {
			return 0
		}
	} else if a.Type == List && b.Type == List {
		return CompareLists(a.List, b.List)
	} else if a.Type == Item {
		return CompareLists([]ListItem{a}, b.List)
	} else if b.Type == Item {
		return CompareLists(a.List, []ListItem{b})
	}

	return -1
}

func CompareLists(a, b []ListItem) int {
	minLen := len(a)
	maxLen := len(b)
	if minLen > maxLen {
		minLen, maxLen = maxLen, minLen
	}

	for i := 0; i < minLen; i++ {
		result := Compare(a[i], b[i])
		if result == -1 {
			continue
		}
		return result
	}

	// ran out of items
	if len(a) < len(b) {
		return 1
	}

	if len(b) < len(a) {
		return 0
	}

	return -1
}

func Run(fileName string) Result {
	lines := utils.ReadFileToLines(fileName)
	var pairs [][]ListItem

	for _, line := range lines {
		if line != "" {
			list := ParseList(line)
			pairs = append(pairs, list)
		}
	}

	sum := 0
	for i := 0; i < len(pairs); i += 2 {
		a := pairs[i]
		b := pairs[i+1]
		result := CompareLists(a, b)
		if result < 0 {
			panic("expected result > 0")
		}
		if result == 1 {
			index := i/2 + 1
			sum += index
		}
	}

	return Result{sum, 0}
}
