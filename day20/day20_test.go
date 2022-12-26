package day20

import (
	"fmt"
	"testing"
)

func TestMove(t *testing.T) {
	fmt.Println("===== TestMove =====")
	list := []int{0, 1, 2, 3, 4}
	fmt.Println(list)

	Move(&list, 1, 3)
	fmt.Println(list)
	Move(&list, 3, 1)
	fmt.Println(list)
	Move(&list, 0, 4)
	fmt.Println(list)
	Move(&list, 4, 0)
	fmt.Println(list)
	Move(&list, 3, 0)
	fmt.Println(list)
	Move(&list, 0, 3)
	fmt.Println(list)
	Move(&list, 4, 2)
	fmt.Println(list)
	Move(&list, 2, 4)
	fmt.Println(list)
}

func TestShift(t *testing.T) {
	fmt.Println("===== TestShift =====")
	list := []int{0, 1, 2, 3, 4}
	fmt.Println(list)

	Shift(&list, 0, 1)
	fmt.Println(list)
	Shift(&list, 1, -1)
	fmt.Println(list)
	Shift(&list, 4, 1)
	fmt.Println(list)
	Shift(&list, 1, 2)
	fmt.Println(list)
	Shift(&list, 3, 1)
	fmt.Println(list)

	Shift(&list, 1, 4)
	fmt.Println(list)

	Shift(&list, 1, -1)
	fmt.Println(list)
	Shift(&list, 4, -1)
	fmt.Println(list)
	Shift(&list, 3, -1)
	fmt.Println(list)
	Shift(&list, 2, -1)
	fmt.Println(list)

	Shift(&list, 1, 5)
	fmt.Println(list)
	Shift(&list, 2, -1)
	fmt.Println(list)

	Shift(&list, 3, -6)
	fmt.Println(list)

	fmt.Println()
	Shift(&list, 2, 2)
	fmt.Println(list)
}

func TestNewIndex(t *testing.T) {
	fmt.Println("===== TestShift =====")
	list := []int{0, 1, 2, 3, 4}

	fmt.Println(list)

	fmt.Println(NewIndex(1, 0, len(list)))
	fmt.Println(NewIndex(2, 0, len(list)))
	fmt.Println(NewIndex(3, 0, len(list)))
	fmt.Println(NewIndex(4, 0, len(list)))
	fmt.Println(NewIndex(8, 0, len(list)))
	fmt.Println(NewIndex(12, 0, len(list)))
	fmt.Println(NewIndex(13, 0, len(list)))
	fmt.Println()

	fmt.Println(NewIndex(3, -1, len(list)))
	fmt.Println(NewIndex(3, -2, len(list)))
	fmt.Println(NewIndex(3, -3, len(list)))
	fmt.Println(NewIndex(3, -4, len(list)))
	fmt.Println(NewIndex(3, -8, len(list)))
	fmt.Println(NewIndex(3, -12, len(list)))
	fmt.Println(NewIndex(3, -13, len(list)))
}

func TestDay20(t *testing.T) {
	fmt.Println("===== Day 20 =====")

	{
		expected := Result{3, 1623178306}
		actual := Run("example.txt")
		if actual != expected {
			t.Errorf("Expected: %d but got: %d\n", expected, actual)
		}
	}

	// {
	// 	expected := Result{7004, 0}
	// 	actual := Run("input.txt")
	// 	if actual != expected {
	// 		t.Errorf("Expected: %d but got: %d\n", expected, actual)
	// 	}
	// }
}
