package day21

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

type Node struct {
	name     string
	value    int
	hasValue bool
	op       string
	lhs      *Node
	rhs      *Node
}

func (n *Node) Eval() int {
	if !n.hasValue {
		lhs := n.lhs.Eval()
		rhs := n.rhs.Eval()

		switch n.op {
		case "+":
			n.value = lhs + rhs
		case "-":
			n.value = lhs - rhs
		case "/":
			n.value = lhs / rhs
		case "*":
			n.value = lhs * rhs
		case "=":
			n.value = 0
			if lhs == rhs {
				n.value = 1
			}
		default:
			panic(fmt.Sprintf("Unknown op: %s", n.op))
		}
	}

	n.hasValue = true
	return n.value
}

func BuildTree(name string, nodes *map[string]Node) *Node {
	node := (*nodes)[name]

	if !node.hasValue {
		node.lhs = BuildTree(node.lhs.name, nodes)
		node.rhs = BuildTree(node.rhs.name, nodes)
	}

	return &node
}

func Run(fileName string) Result {
	lines := utils.ReadFileToLines(fileName)

	nodes := map[string]Node{}
	for _, line := range lines {
		fields := strings.Split(line, ": ")

		node := Node{}
		node.name = fields[0]

		value, err := strconv.Atoi(fields[1])
		if err != nil {
			fields = strings.Fields(fields[1])
			node.op = fields[1]
			node.lhs = &Node{}
			node.rhs = &Node{}
			node.lhs.name = fields[0]
			node.rhs.name = fields[2]
		} else {
			node.value = value
			node.hasValue = true
		}

		nodes[node.name] = node
	}

	root := BuildTree("root", &nodes)
	result := root.Eval()

	return Result{result, 0}
}
