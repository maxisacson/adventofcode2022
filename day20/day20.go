package day20

import (
	"aoc22/utils"
	"fmt"
	"strconv"
)

type Result struct {
	part1 int
	part2 int
}

func AppendList(a, b []int) []int {
	for _, x := range b {
		a = append(a, x)
	}

	return a
}

func Insert(list []int, index, value int) []int {
	list = append(list, 0)
	copy(list[index+1:], list[index:])
	list[index] = value

	return list
}

func MixOnce(numbers []int, order []int) ([]int, []int) {

	index := order[len(order)-1]
	order = order[:len(order)-1]

	newIndex := (index + numbers[index]) % len(numbers)
	for newIndex < 0 {
		newIndex += len(numbers)
	}
	fmt.Printf("%d -> %d\n", index, newIndex)

	left := make([]int, index)
	right := make([]int, len(numbers)-index-1)
	copy(left, numbers[:index])
	copy(right, numbers[index+1:])

	if newIndex < len(left) {
		left = Insert(left, newIndex, numbers[index])
	} else {
		rightIndex := newIndex - len(left)
		right = Insert(right, rightIndex, numbers[index])
	}

	left = append(left, right...)

	for i := len(order) - 1; i >= 0; i-- {
		if order[i] <= newIndex-index {
			order[i]--
		}
	}

	return left, order
}

func Run(fileName string) Result {
	lines := utils.ReadFileToLines(fileName)

	numbers := make([]int, len(lines))
	order := make([]int, len(lines))
	for i, line := range lines {
		numbers[i], _ = strconv.Atoi(line)
		order[i] = len(lines) - 1 - i
	}

	fmt.Println(numbers)
	fmt.Println(order)
	fmt.Println()

	numbers, order = MixOnce(numbers, order)
	fmt.Println(numbers)
	fmt.Println(order)
	fmt.Println()

	numbers, order = MixOnce(numbers, order)
	fmt.Println(numbers)
	fmt.Println(order)
	fmt.Println()

	numbers, order = MixOnce(numbers, order)
	fmt.Println(numbers)
	fmt.Println(order)
	fmt.Println()

	numbers, order = MixOnce(numbers, order)
	fmt.Println(numbers)
	fmt.Println(order)
	fmt.Println()

	return Result{len(lines), len(lines)}
}
