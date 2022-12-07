package day7

import (
	"fmt"
	"testing"
)

func TestDay7(t *testing.T) {
	fmt.Println("===== Day 7 =====")

	{
		expected := Result{95437, 24933642}
		actual := Run("example.txt")
		if actual != expected {
			t.Errorf("Expected: %d but got: %d\n", expected, actual)
		}
	}

	{
		expected := Result{919137, 2877389}
		actual := Run("input.txt")
		if actual != expected {
			t.Errorf("Expected: %d but got: %d\n", expected, actual)
		}
		fmt.Println("--- Part 1 ---")
		fmt.Println("Sum:", actual.part1)
		fmt.Println("--- Part 2 ---")
		fmt.Println("Size:", actual.part2)
	}
}
