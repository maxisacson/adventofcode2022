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

func Move(list *[]int, oldIndex, newIndex int) {
	N := len(*list)
	left := make([]int, oldIndex)
	right := make([]int, N-oldIndex-1)
	copy(left, (*list)[:oldIndex])
	copy(right, (*list)[oldIndex+1:])

	if newIndex < len(left) {
		left = Insert(left, newIndex, (*list)[oldIndex])
	} else {
		rightIndex := newIndex - len(left)
		right = Insert(right, rightIndex, (*list)[oldIndex])
	}

	*list = append(left, right...)
}

func MixOnce(numbers, stack []int) ([]int, []int) {

	index := Pop(&stack)
	N := len(numbers)

	newIndex := (index + numbers[index]) % (N - 1)
	for newIndex < 0 {
		newIndex += (N - 1)
	}
	if newIndex == 0 && numbers[index] < 0 {
		newIndex = N - 1
	}

	Move(&numbers, index, newIndex)

	for i := 0; i < len(stack); i++ {
		if index < stack[i] && stack[i] <= newIndex {
			stack[i]--
		}
	}

	return numbers, stack
}

func Run(fileName string) Result {
	lines := utils.ReadFileToLines(fileName)

	numbers := make([]int, len(lines))
	order := make([]int, len(lines))
	for i, line := range lines {
		numbers[i], _ = strconv.Atoi(line)
		order[i] = len(lines) - 1 - i
	}

	// part 1
	numbers1 := make([]int, len(numbers))
	order1 := make([]int, len(order))
	copy(numbers1, numbers)
	copy(order1, order)
	for len(order1) > 0 {
		numbers1, order1 = MixOnce(numbers1, order1)
	}
	for _, n := range numbers {
		fmt.Print(n, ", ")
	}
	fmt.Println()
	for _, val := range order1 {
		fmt.Print(numbers1[val], ", ")
	}
	fmt.Println()

	index := 0
	for i, val := range numbers1 {
		if val == 0 {
			index = i
			break
		}
	}

	sum := numbers1[(index+1000)%len(numbers1)]
	sum += numbers1[(index+2000)%len(numbers1)]
	sum += numbers1[(index+3000)%len(numbers1)]

	// part 2
	key := 811589153
	// key := 1
	numbers2 := make([]int, len(numbers))
	orderMap := map[int]int{}
	for i, val := range numbers {
		newVal := key * val
		numbers2[i] = newVal
		orderMap[newVal] = i
	}
	fmt.Println(orderMap)

	// fmt.Println(numbers2)
	fmt.Println()
	for round := 0; round < 10; round++ {
		order2 := make([]int, len(numbers2))
		for i, val := range numbers2 {
			j := orderMap[val]
			order2[len(order2)-1-j] = i
		}
		fmt.Println(order2)
		for len(order2) > 0 {
			numbers2, order2 = MixOnce(numbers2, order2)
		}
		// fmt.Println(numbers2)
		// fmt.Println(order2)
		// for _, n := range numbers {
		// 	fmt.Print(n, ", ")
		// }
		// fmt.Println()
		// for _, val := range order2 {
		// 	// j := len(order2) - i - 1
		// 	fmt.Print(numbers2[val], ", ")
		//
		// 	// if numbers[j] != numbers2[val] {
		// 	// 	panic("help!")
		// 	// }
		// 	// if numbers[len(order2)-i-1] != numbers2[val] {
		// 	// 	panic("help!")
		// 	// }
		// }
		// fmt.Println()
		fmt.Println(numbers2)
		fmt.Println()
	}

	index = 0
	for i, val := range numbers2 {
		if val == 0 {
			index = i
			break
		}
	}
	// fmt.Println(numbers2[(index+1000)%len(numbers2)])
	// fmt.Println(numbers2[(index+2000)%len(numbers2)])
	// fmt.Println(numbers2[(index+3000)%len(numbers2)])

	sum2 := numbers2[(index+1000)%len(numbers2)]
	sum2 += numbers2[(index+2000)%len(numbers2)]
	sum2 += numbers2[(index+3000)%len(numbers2)]

	return Result{sum, sum2}
}
