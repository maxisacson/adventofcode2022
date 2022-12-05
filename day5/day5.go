package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func push(stack *[]string, element string) {
	*stack = append(*stack, element)
}

func pushN(stack *[]string, element []string) {
	for _, e := range element {
		push(stack, e)
	}
}

func pop(stack *[]string) string {
	n := len(*stack) - 1
	element := (*stack)[n]
	*stack = (*stack)[:n]
	return element
}

func popN(stack *[]string, count int) []string {
	var result []string
	for i := 0; i < count; i++ {
		push(&result, pop(stack))
	}
	return result
}

func popNR(stack *[]string, count int) []string {
	var result = make([]string, count)
	for i := 0; i < count; i++ {
		result[count-i-1] = pop(stack)
	}
	return result
}

func atoi(s string) int {
	i, e := strconv.Atoi(s)
	if e != nil {
		panic(e)
	}
	return i
}

func peek(stack *[]string) string {
	return (*stack)[len(*stack)-1]
}

func buildStacks(crateMap *[]string) [][]string {
	numbers := strings.Fields((*crateMap)[len(*crateMap)-1])
	numStacks := len(numbers)

	stacks := make([][]string, numStacks)

	for i := len(*crateMap) - 2; i >= 0; i-- {
		crates := (*crateMap)[i]
		for j := range numbers {
			k := 1 + 4*j
			if k >= len(crates) {
				break
			}
			if string(crates[k]) != " " {
				push(&stacks[j], string(crates[k]))
			}

		}
	}

	return stacks
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	var crateMap []string
	var instructions []string
	parseInstructions := false
	for _, line := range lines {
		if line == "" {
			parseInstructions = true
			continue
		}

		if parseInstructions {
			push(&instructions, line)
		} else {
			push(&crateMap, line)
		}
	}

	stacks := buildStacks(&crateMap)

	for _, instruction := range instructions {
		fields := strings.Fields(instruction)
		count := atoi(fields[1])
		from := atoi(fields[3]) - 1
		to := atoi(fields[5]) - 1

		crates := popN(&stacks[from], count)
		pushN(&stacks[to], crates)
	}

	message := ""
	for _, stack := range stacks {
		message = message + peek(&stack)
	}

	if message != "VRWBSFZWM" && os.Args[1] == "input.txt" {
		panic("wrong answer!")
	}

	fmt.Println("--- Part 1 ---")
	fmt.Println("Message:", message)

	stacks = buildStacks(&crateMap)

	for _, instruction := range instructions {
		fields := strings.Fields(instruction)
		count := atoi(fields[1])
		from := atoi(fields[3]) - 1
		to := atoi(fields[5]) - 1

		crates := popNR(&stacks[from], count)
		pushN(&stacks[to], crates)
	}

	message = ""
	for _, stack := range stacks {
		message = message + peek(&stack)
	}

	if message != "RBTWJWMCF" && os.Args[1] == "input.txt" {
		panic("wrong answer!")
	}

	fmt.Println("--- Part 2 ---")
	fmt.Println("Message:", message)
}
