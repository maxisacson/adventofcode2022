package day22

import (
	"fmt"
	"testing"
)

func TestDay22(t *testing.T) {
	fmt.Println("===== Day 22 =====")

	{
		expected := Result{6032, 5031}
		actual := Run("example.txt")
		if actual != expected {
			t.Errorf("Expected: %d but got: %d\n", expected, actual)
		}
	}

	{
		expected := Result{181128, 0}
		actual := Run("input.txt")
		if actual != expected {
			t.Errorf("Expected: %d but got: %d\n", expected, actual)
		}
	}
}
