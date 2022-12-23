package day23

import (
	"fmt"
	"testing"
)

func TestDay23(t *testing.T) {
	fmt.Println("===== Day 23 =====")

	{
		expected := Result{110, 20}
		actual := Run("example.txt")
		if actual != expected {
			t.Errorf("Expected: %d but got: %d\n", expected, actual)
		}
	}

	{
		expected := Result{4068, 968}
		actual := Run("input.txt")
		if actual != expected {
			t.Errorf("Expected: %d but got: %d\n", expected, actual)
		}
	}
}
