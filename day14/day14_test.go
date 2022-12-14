package day14

import (
	"fmt"
	"testing"
)

func TestDay14(t *testing.T) {
	fmt.Println("===== Day 14 =====")

	{
		expected := Result{24, 93}
		actual := Run("example.txt")
		if actual != expected {
			t.Errorf("Expected: %d but got: %d\n", expected, actual)
		}
	}

	{
		expected := Result{0, 0}
		actual := Run("input.txt")
		if actual != expected {
			t.Errorf("Expected: %d but got: %d\n", expected, actual)
		}
	}
}
