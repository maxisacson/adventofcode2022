package day3

import (
	"fmt"
	"testing"
)

func TestDay3(t *testing.T) {
	fmt.Println("===== Day 3 =====")

	{
		expected := Result{157, 70}
		actual := Run("example.txt")
		if actual != expected {
			t.Errorf("Expected: %d but got: %d\n", expected, actual)
		}
	}

	{
		expected := Result{7597, 2607}
		actual := Run("input.txt")
		if actual != expected {
			t.Errorf("Expected: %d but got: %d\n", expected, actual)
		}
	}
}
