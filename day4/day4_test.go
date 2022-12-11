package day4

import (
	"fmt"
	"testing"
)

func TestDay4(t *testing.T) {
	fmt.Println("===== Day 4 =====")

	{
		expected := Result{2, 4}
		actual := Run("example.txt")
		if actual != expected {
			t.Errorf("Expected: %d but got: %d\n", expected, actual)
		}
	}

	{
		expected := Result{532, 854}
		actual := Run("input.txt")
		if actual != expected {
			t.Errorf("Expected: %d but got: %d\n", expected, actual)
		}
	}
}
