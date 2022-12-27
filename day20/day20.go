package day20

import (
	"aoc22/utils"
	"strconv"
)

type Result struct {
	part1 int
	part2 int
}

func Insert(list []int, index, value int) []int {
	list = append(list, 0)
	copy(list[index+1:], list[index:])
	list[index] = value

	return list
}

func Pop(list *[]int) int {
	n := len(*list) - 1
	x := (*list)[n]
	*list = (*list)[:n]

	return x
}

func NewIndex(shift, index, length int) int {
	if shift%(length-1) == 0 {
		return index
	}

	if index+shift <= 0 {
		return ((index + shift) % (length - 1)) + (length - 1)
	}

	return (index + shift) % (length - 1)
}

func Shift(list *[]int, index, shift int) int {
	length := len(*list)
	newIndex := NewIndex(shift, index, length)

	Move(list, index, newIndex)

	return newIndex
}

func Move(list *[]int, index, newIndex int) {
	if newIndex < 0 || newIndex >= len(*list) {
		panic("newIndex out of bounds")
	}
	value := (*list)[index]
	copy((*list)[index:], (*list)[index+1:])
	copy((*list)[newIndex+1:], (*list)[newIndex:])
	(*list)[newIndex] = value
}

func MixOnce(numbers, stack *[]int, order *[]int) {
	index := Pop(stack)

	n := (*numbers)[index]
	newIndex := Shift(numbers, index, n)

	UpdateStack(stack, index, newIndex)

	if order != nil {
		for i, x := range *order {
			if index <= x && x <= newIndex {
				(*order)[i]--
			} else if newIndex <= x && x <= index {
				(*order)[i]++
			}
		}
		j := len(*order) - len(*stack) - 1
		(*order)[j] = newIndex
	}
}

func UpdateStack(stack *[]int, index, newIndex int) {
	for i, x := range *stack {
		if index < x && x <= newIndex {
			(*stack)[i]--
		} else if newIndex <= x && x < index {
			(*stack)[i]++
		}
	}
}

func IndexOf(value int, list *[]int) int {
	for i, n := range *list {
		if n == value {
			return i
		}
	}

	return len(*list)
}

func Reversed(list *[]int) []int {
	N := len(*list)
	reversed := make([]int, N)
	for i, x := range *list {
		reversed[N-1-i] = x
	}
	return reversed
}

func MakeStack(order *[]int) []int {
	return Reversed(order)
}

func Run(fileName string) Result {
	lines := utils.ReadFileToLines(fileName)

	numbers := make([]int, len(lines))
	order := make([]int, len(lines))
	for i, line := range lines {
		numbers[i], _ = strconv.Atoi(line)
		order[i] = i
	}

	// part 1
	stack := MakeStack(&order)
	for len(stack) > 0 {
		MixOnce(&numbers, &stack, &order)
	}

	ind := IndexOf(0, &numbers)
	N := len(numbers)
	sum := numbers[(ind+1000)%N]
	sum += numbers[(ind+2000)%N]
	sum += numbers[(ind+3000)%N]

	// part 2
	for i, line := range lines {
		numbers[i], _ = strconv.Atoi(line)
		numbers[i] *= 811589153
		order[i] = i
	}

	for round := 0; round < 10; round++ {
		stack = MakeStack(&order)
		for len(stack) > 0 {
			MixOnce(&numbers, &stack, &order)
		}
		stack = MakeStack(&order)
	}

	ind = IndexOf(0, &numbers)
	sum2 := numbers[(ind+1000)%N]
	sum2 += numbers[(ind+2000)%N]
	sum2 += numbers[(ind+3000)%N]

	return Result{sum, sum2}
}
