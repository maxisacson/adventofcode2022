package day12

import (
	"fmt"
	"testing"
)

func TestDay12(t *testing.T) {
	fmt.Println("===== Day 12 =====")

	{
		expected := Result{31, 29}
		actual := Run("example.txt")
		if actual != expected {
			t.Errorf("Expected: %d but got: %d\n", expected, actual)
		}
	}

	{
		expected := Result{437, 430}
		actual := Run("input.txt")
		if actual != expected {
			t.Errorf("Expected: %d but got: %d\n", expected, actual)
		}
	}
}
