package day6

import (
	"fmt"
	"testing"
)

func TestDay6(t *testing.T) {
	fmt.Println("===== Day 6 =====")

	{
		expected := Result{11, 26}
		actual := Run("example.txt")
		if actual != expected {
			t.Errorf("Expected: %d but got: %d\n", expected, actual)
		}
	}

	{
		expected := Result{1929, 3298}
		actual := Run("input.txt")
		if actual != expected {
			t.Errorf("Expected: %d but got: %d\n", expected, actual)
		}
	}
}
