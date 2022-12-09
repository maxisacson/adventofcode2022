package day8

import (
	"fmt"
	"testing"
)

func TestDay8(t *testing.T) {
	fmt.Println("===== Day 8 =====")

	{
		expected := Result{21, 8}
		actual := Run("example.txt")
		if actual != expected {
			t.Errorf("Expected: %d but got: %d\n", expected, actual)
		}
	}

	{
		expected := Result{1870, 517440}
		actual := Run("input.txt")
		if actual != expected {
			t.Errorf("Expected: %d but got: %d\n", expected, actual)
		}
	}
}
