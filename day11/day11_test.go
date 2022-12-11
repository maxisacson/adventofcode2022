package day11

import (
	"fmt"
	"testing"
)

func TestDay11(t *testing.T) {
	fmt.Println("===== Day 11 =====")

	{
		expected := Result{10605, 0}
		actual := Run("example.txt")
		if actual != expected {
			t.Errorf("Expected: %d but got: %d\n", expected, actual)
		}
	}

	{
		expected := Result{57838, 0}
		actual := Run("input.txt")
		if actual != expected {
			t.Errorf("Expected: %d but got: %d\n", expected, actual)
		}
	}
}
