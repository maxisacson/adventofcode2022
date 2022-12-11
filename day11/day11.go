package day11

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

type Monkey struct {
	items []int
	op    []string
	test  int
	next  []int
	count int
}

func (m *Monkey) DoRound(monkies *[]Monkey) {
	for _, item := range m.items {
		value := m.DoOp(item) / 3

		var next int
		if value%m.test == 0 {
			next = m.next[0]
		} else {
			next = m.next[1]
		}

		(*monkies)[next].Catch(value)

		m.count++
	}

	m.items = []int{}
}

func (m *Monkey) Catch(item int) {
	m.items = append(m.items, item)
}

func (m *Monkey) DoOp(item int) int {
	var value1 int
	var value2 int

	if m.op[0] == "old" {
		value1 = item
	} else {
		value1, _ = strconv.Atoi(m.op[0])
	}

	if m.op[2] == "old" {
		value2 = item
	} else {
		value2, _ = strconv.Atoi(m.op[2])
	}

	switch m.op[1] {
	case "+":
		return value1 + value2
	case "*":
		return value1 * value2
	default:
		panic(fmt.Sprintf("Unknown op: %s", m.op[1]))
	}
}

func Run(fileName string) Result {
	lines := utils.ReadFileToLines(fileName)

	monkies := []Monkey{}

	for i := 0; i < len(lines); i += 7 {
		monkey := Monkey{}
		fields := strings.Split(lines[i+1], ": ")
		items := strings.Split(fields[1], ", ")
		for _, item := range items {
			it, _ := strconv.Atoi(item)
			monkey.items = append(monkey.items, it)
		}

		fields = strings.Split(lines[i+2], " = ")
		monkey.op = strings.Fields(fields[1])

		fields = strings.Fields(lines[i+3])
		test, _ := strconv.Atoi(fields[len(fields)-1])
		monkey.test = test

		fields = strings.Fields(lines[i+4])
		next0, _ := strconv.Atoi(fields[len(fields)-1])

		fields = strings.Fields(lines[i+5])
		next1, _ := strconv.Atoi(fields[len(fields)-1])

		monkey.next = []int{next0, next1}

		monkies = append(monkies, monkey)
	}

	for i := 0; i < 20; i++ {
		for j := 0; j < len(monkies); j++ {
			monkies[j].DoRound(&monkies)
		}
	}

	sort.Slice(monkies, func(i, j int) bool {
		return monkies[i].count > monkies[j].count
	})

	monkeyBusiness := monkies[0].count * monkies[1].count

	return Result{monkeyBusiness, 0}
}
