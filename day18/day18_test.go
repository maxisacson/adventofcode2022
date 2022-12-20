package day18

import (
	"fmt"
	"testing"
)

func TestDay18(t *testing.T) {
	fmt.Println("===== Day 18 =====")

	{
		expected := Result{64, 0}
		actual := Run("example.txt")
		if actual != expected {
			t.Errorf("Expected: %d but got: %d\n", expected, actual)
		}
	}

	{
		expected := Result{4288, 0}
		actual := Run("input.txt")
		if actual != expected {
			t.Errorf("Expected: %d but got: %d\n", expected, actual)
		}
	}
}
