package day15

import (
	"fmt"
	"testing"
)

func TestDay15(t *testing.T) {
	fmt.Println("===== Day 15 =====")

	{
		expected := Result{26, 0}
		actual := Run("example.txt", 10)
		if actual != expected {
			t.Errorf("Expected: %d but got: %d\n", expected, actual)
		}
	}

	{
		expected := Result{4985193, 0}
		actual := Run("input.txt", 2000000)
		if actual != expected {
			t.Errorf("Expected: %d but got: %d\n", expected, actual)
		}
	}
}
