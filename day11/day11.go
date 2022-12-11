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
	norm  int
	gcd   int
}

func (m *Monkey) DoRound(monkies *[]Monkey) {
	for _, item := range m.items {
		value := m.DoOp(item) / m.norm
		if m.gcd > 0 {
			value %= m.gcd
		}

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

func NewMonkey(i int, lines []string, norm int) Monkey {
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

	monkey.norm = norm
	monkey.gcd = 0

	return monkey
}

func Run(fileName string) Result {
	lines := utils.ReadFileToLines(fileName)

	monkies := []Monkey{}

	// Part 1
	for i := 0; i < len(lines); i += 7 {
		monkies = append(monkies, NewMonkey(i, lines, 3))
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

	// Part 2
	monkies = []Monkey{}
	gcd := 1
	for i := 0; i < len(lines); i += 7 {
		monkies = append(monkies, NewMonkey(i, lines, 1))
		gcd *= monkies[len(monkies)-1].test
	}
	for j := 0; j < len(monkies); j++ {
		monkies[j].gcd = gcd
	}

	for i := 0; i < 10000; i++ {
		for j := 0; j < len(monkies); j++ {
			monkies[j].DoRound(&monkies)
		}
	}

	sort.Slice(monkies, func(i, j int) bool {
		return monkies[i].count > monkies[j].count
	})

	monkeyBusiness2 := monkies[0].count * monkies[1].count

	return Result{monkeyBusiness, monkeyBusiness2}
}
