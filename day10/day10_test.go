package day10

import (
	"fmt"
	"testing"
)

func TestDay10(t *testing.T) {
	fmt.Println("===== Day 10 =====")

	{
		expected := Result{13140, 0}
		actual := Run("example.txt")
		if actual != expected {
			t.Errorf("Expected: %d but got: %d\n", expected, actual)
		}
	}

	{
		expected := Result{17940, 0}
		actual := Run("input.txt")
		if actual != expected {
			t.Errorf("Expected: %d but got: %d\n", expected, actual)
		}
	}
}
