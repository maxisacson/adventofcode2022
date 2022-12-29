package day16

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

type Valve struct {
	name     string
	flowRate int
	tunnels  []string
	isOpen   bool
}

type Edge struct {
	to   string
	from string
	cost int
}

type Node struct {
	label string
	value int
	edges []Edge
}

type Graph struct {
	nodes map[string]Node
}

func AllValvesOpen(valves map[string]Valve) bool {
	for _, v := range valves {
		if v.flowRate > 0 && !v.isOpen {
			return false
		}
	}

	return true
}

func copyValves(valves map[string]Valve) map[string]Valve {
	ret := make(map[string]Valve)
	for k, v := range valves {
		ret[k] = v
	}
	return ret
}

func copyInto(dst *map[string]Valve, src map[string]Valve) {
	for k, v := range src {
		(*dst)[k] = v
	}
}

func Traverse(graph Graph, valves map[string]Valve, name string, previous string, timeLeft int, path *[]string) int {
	if timeLeft <= 1 {
		return 0
	}

	if AllValvesOpen(valves) {
		return 0
	}

	valve := valves[name]
	valvesOpen := copyValves(valves)
	valvesClosed := copyValves(valves)
	*path = append(*path, fmt.Sprintf("%s:%d", name, timeLeft))
	newPath := []string{}

	node := graph.nodes[name]

	maxPressure := -1
	if !valve.isOpen && valve.flowRate > 0 {
		pressure := valve.flowRate * (timeLeft - 1)
		valve.isOpen = true
		valvesOpen[name] = valve

		for _, edge := range node.edges {
			timeLeftNext := timeLeft - 1 - edge.cost
			if timeLeftNext < 1 {
				continue
			}

			next := edge.to
			valvesCopy := copyValves(valvesOpen)

			pathCopy := make([]string, len(*path))
			copy(pathCopy, *path)
			pathCopy = append(pathCopy, fmt.Sprintf("%s:%d", name, timeLeft-1))

			thisPressure := pressure + Traverse(graph, valvesCopy, next, name, timeLeftNext, &pathCopy)
			if thisPressure > maxPressure {
				maxPressure = thisPressure
				copyInto(&valves, valvesCopy)
				newPath = pathCopy
			}
		}
	}

	for _, edge := range node.edges {
		next := edge.to
		if next == previous {
			continue
		}
		timeLeftNext := timeLeft - edge.cost
		if timeLeftNext < 1 {
			continue
		}
		valvesCopy := copyValves(valvesClosed)

		pathCopy := make([]string, len(*path))
		copy(pathCopy, *path)

		thisPressure := Traverse(graph, valvesCopy, next, name, timeLeftNext, &pathCopy)
		if thisPressure > maxPressure {
			maxPressure = thisPressure
			copyInto(&valves, valvesCopy)
			newPath = pathCopy
		}
	}

	// for _, v := range valves {
	// 	fmt.Println(v.name, v.isOpen)
	// }
	// fmt.Println()
	*path = newPath
	return maxPressure
}

func BuildGraph(valves map[string]Valve) Graph {
	graph := Graph{}
	graph.nodes = make(map[string]Node)
	for k, v := range valves {
		node := Node{}
		node.label = k
		node.value = v.flowRate
		for _, next := range v.tunnels {
			edge := Edge{}
			edge.to = next
			edge.from = node.label
			edge.cost = 1
			node.edges = append(node.edges, edge)
		}
		graph.nodes[node.label] = node
	}
	return graph
}

func (g *Graph) Contract() {
	newNodes := make(map[string]Node)
	for _, node := range g.nodes {
		edges := []Edge{}
		for _, edge := range node.edges {
			contracted := g.ContractEdge(edge)
			edges = append(edges, contracted...)
		}
		node.edges = edges
		newNodes[node.label] = node
	}
	g.nodes = newNodes
}

func (g *Graph) ContractEdge(edge Edge) []Edge {
	contracted := []Edge{}
	if g.nodes[edge.to].value > 0 {
		contracted = append(contracted, edge)
		return contracted
	}

	to := g.nodes[edge.to]
	for _, e := range to.edges {
		if e.to == edge.from {
			continue
		}
		tmp := g.ContractEdge(e)
		for i := range tmp {
			tmp[i].from = edge.from
			tmp[i].cost += edge.cost
		}
		contracted = append(contracted, tmp...)
	}

	return contracted
}

func (g Graph) String() string {
	lines := []string{}
	for _, node := range g.nodes {
		for _, edge := range node.edges {
			to := g.nodes[edge.to]
			lines = append(lines, fmt.Sprintf("%s(%d) -{%d}-> %s(%d)", node.label, node.value, edge.cost, to.label, to.value))
		}
	}

	return strings.Join(lines, "\n")
}

func Run(fileName string) Result {
	lines := utils.ReadFileToLines(fileName)

	valves := make(map[string]Valve)

	for _, line := range lines {
		parts := strings.Split(line, "; ")
		fields0 := strings.Fields(parts[0])
		valve := fields0[1]
		flowRate, _ := strconv.Atoi(strings.Split(fields0[4], "=")[1])
		var connections []string
		if parts[1][16:22] == "valves" {
			connections = strings.Split(parts[1][23:len(parts[1])], ", ")
		} else {
			connections = []string{parts[1][22:len(parts[1])]}
		}
		valves[valve] = Valve{valve, flowRate, connections, false}
		// fmt.Println(valves[valve])
	}
	// fmt.Println()

	graph := BuildGraph(valves)
	// fmt.Println(graph)
	// fmt.Println()
	graph.Contract()
	// fmt.Println(graph)
	// fmt.Println()

	pressure := 0
	path := []string{}
	// pressure = Traverse(graph, valves, "AA", "", 30, &path)
	pressure = Traverse(graph, valves, "AA", "", 30, &path)

	fmt.Println(path)
	fmt.Println(pressure)
	// for _, v := range valves {
	// 	fmt.Println(v.name, v.isOpen, v.flowRate)
	// }
	// fmt.Println()

	return Result{pressure, 0}
}
