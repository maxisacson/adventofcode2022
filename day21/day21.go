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

	isVariable bool
}

func (n *Node) Eval() {
	if n.hasValue {
		return
	}

	if n.isVariable {
		return
	}

	n.lhs.Eval()
	n.rhs.Eval()

	if n.lhs.hasValue && n.rhs.hasValue {
		lhs := n.lhs.value
		rhs := n.rhs.value

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
		n.hasValue = true
	}
}

func (n *Node) Solve() {
	for !n.lhs.isVariable && !n.rhs.isVariable {
		n.Eval()

		if n.lhs.hasValue && n.rhs.hasValue {
			return
		}

		if !n.lhs.hasValue && !n.rhs.hasValue {
			PrintTree(n, 0)
			panic("At least one subtree needs to be known!")
		}

		if !n.lhs.hasValue {
			invOp := ""
			switch n.lhs.op {
			case "+":
				invOp = "-"
			case "-":
				invOp = "+"
			case "/":
				invOp = "*"
			case "*":
				invOp = "/"
			}

			node := new(Node)
			node.name = "tmp"
			node.op = invOp

			if !n.lhs.rhs.hasValue {
				if n.lhs.op == "+" || n.lhs.op == "*" {
					tmp := n.lhs.rhs
					n.lhs.rhs = n.lhs.lhs
					n.lhs.lhs = tmp
				}
			}

			node.rhs = n.lhs.rhs
			node.lhs = n.rhs
			n.rhs = node
			n.lhs = n.lhs.lhs
		} else {
			tmp := n.lhs
			n.lhs = n.rhs
			n.rhs = tmp
		}
	}

	n.Eval()
	if !n.lhs.isVariable {
		tmp := n.lhs
		n.lhs = n.rhs
		n.rhs = tmp
	}
}

func BuildTree(name string, nodes *map[string]Node) *Node {
	node := (*nodes)[name]

	if !node.hasValue && !node.isVariable {
		node.lhs = BuildTree(node.lhs.name, nodes)
		node.rhs = BuildTree(node.rhs.name, nodes)
	}

	return &node
}

func PrintTree(root *Node, indentLevel int) {
	indent := strings.Repeat(" ", indentLevel)
	if root.hasValue {
		fmt.Printf("%s%s: %d\n", indent, root.name, root.value)
	} else if root.isVariable {
		fmt.Printf("%s%s: X\n", indent, root.name)
	} else {
		fmt.Printf("%s%s: %s [\n", indent, root.name, root.op)
		PrintTree(root.lhs, indentLevel+2)
		PrintTree(root.rhs, indentLevel+2)
		fmt.Printf("%s]\n", indent)
	}
}

func (n Node) String() string {
	if n.isVariable {
		return "X"
	} else if n.hasValue {
		return fmt.Sprint(n.value)
	} else if n.op == "=" {
		return fmt.Sprintf("%v %s %v", n.lhs, n.op, n.rhs)
	}

	return fmt.Sprintf("(%v%s%v)", n.lhs, n.op, n.rhs)
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

	nodesCopy := nodes

	// Part 1
	root := BuildTree("root", &nodes)
	root.Eval()
	result := root.value

	// Part 2
	root2 := nodesCopy["root"]
	root2.op = "="
	nodesCopy["root"] = root2
	humn := nodesCopy["humn"]
	humn.hasValue = false
	humn.isVariable = true
	nodesCopy["humn"] = humn

	root = BuildTree("root", &nodesCopy)
	root.Solve()

	return Result{result, root.rhs.value}
}
