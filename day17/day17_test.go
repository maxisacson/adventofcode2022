package day17

import (
	"fmt"
	"testing"
)

func TestDay17(t *testing.T) {
	fmt.Println("===== Day 17 =====")

	{
		expected := Result{3068, 1514285714288}
		actual := Run("example.txt")
		if actual != expected {
			t.Errorf("Expected: %d but got: %d\n", expected, actual)
		}
	}

	{
		expected := Result{3186, 1566376811584}
		actual := Run("input.txt")
		if actual != expected {
			t.Errorf("Expected: %d but got: %d\n", expected, actual)
		}
	}
}
