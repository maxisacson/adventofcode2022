package day24

import (
	"fmt"
	"testing"
)

func TestDay24(t *testing.T) {
	fmt.Println("===== Day 24 =====")

	{
		expected := Result{18, 54}
		actual := Run("example.txt")
		if actual != expected {
			t.Errorf("Expected: %d but got: %d\n", expected, actual)
		}
	}

	{
		expected := Result{251, 758}
		actual := Run("input.txt")
		if actual != expected {
			t.Errorf("Expected: %d but got: %d\n", expected, actual)
		}
	}
}
