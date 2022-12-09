package day9

import (
	"fmt"
	"testing"
)

func TestDay9(t *testing.T) {
	fmt.Println("===== Day 9 =====")

	{
		expected := Result{13, 1}
		actual := Run("example.txt")
		if actual != expected {
			t.Errorf("Expected: %d but got: %d\n", expected, actual)
		}
	}

	{
		expected := 36
		actual := Run("example2.txt")
		if actual.part2 != expected {
			t.Errorf("Expected: %d but got: %d\n", expected, actual)
		}
	}

	{
		expected := Result{6256, 2665}
		actual := Run("input.txt")
		if actual != expected {
			t.Errorf("Expected: %d but got: %d\n", expected, actual)
		}
	}
}
