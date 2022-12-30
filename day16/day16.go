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

func Traverse(graph Graph, valves map[string]Valve, name string, previous string, timeLeft int) int {
	if timeLeft <= 1 {
		return 0
	}

	if AllValvesOpen(valves) {
		return 0
	}

	valve := valves[name]
	node := graph.nodes[name]

	maxPressure := -1
	if !valve.isOpen && valve.flowRate > 0 {
		pressure := valve.flowRate * (timeLeft - 1)

		for _, edge := range node.edges {
			timeLeftNext := timeLeft - 1 - edge.cost

			valve.isOpen = true
			valves[name] = valve

			thisPressure := pressure + Traverse(graph, valves, edge.to, name, timeLeftNext)

			valve.isOpen = false
			valves[name] = valve

			if thisPressure > maxPressure {
				maxPressure = thisPressure
			}

		}
	}

	for _, edge := range node.edges {
		if edge.to == previous {
			continue
		}
		timeLeftNext := timeLeft - edge.cost

		thisPressure := Traverse(graph, valves, edge.to, name, timeLeftNext)
		if thisPressure > maxPressure {
			maxPressure = thisPressure
		}
	}

	return maxPressure
}

func TraverseDouble(graph Graph, valves map[string]Valve, name1, name2, prev1, prev2 string, timeLeft1, timeLeft2 int) int {
	if timeLeft1 <= 1 && timeLeft2 <= 1 {
		return 0
	}

	if AllValvesOpen(valves) {
		return 0
	}

	valve1 := valves[name1]

	node1 := graph.nodes[name1]

	maxPressure := -1

	if !valve1.isOpen && valve1.flowRate > 0 {
		pressure := valve1.flowRate * (timeLeft1 - 1)

		for _, edge1 := range node1.edges {
			timeLeftNext1 := timeLeft1 - 1 - edge1.cost

			valve1.isOpen = true
			valves[name1] = valve1

			valve2 := valves[name2]
			node2 := graph.nodes[name2]

			if !valve2.isOpen && valve2.flowRate > 0 {

				for _, edge2 := range node2.edges {
					timeLeftNext2 := timeLeft2 - 1 - edge2.cost
					if timeLeftNext2 < 1 {
						continue
					}

					valve2.isOpen = true
					valves[name2] = valve2

					thisPressure := pressure + TraverseDouble(graph, valves, edge1.to, edge2.to, name1, name2, timeLeftNext1, timeLeftNext2)

					valve2.isOpen = false
					valves[name2] = valve2

					if thisPressure > maxPressure {
						maxPressure = thisPressure
					}
				}
			}

			valve1.isOpen = false
			valves[name1] = valve1

		}
	}

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
	// pressure = Traverse(graph, valves, "AA", "", 25, &path)
	pressure = Traverse(graph, valves, "AA", "", 30)

	// fmt.Println(path)
	// fmt.Println(pressure)
	// for _, v := range valves {
	// 	fmt.Println(v.name, v.isOpen, v.flowRate)
	// }
	// fmt.Println()

	pressure2 := TraverseDouble(graph, valves, "AA", "AA", "", "", 26, 26)

	return Result{pressure, pressure2}
}
