package day1

import (
	"fmt"
	"testing"
)

func TestDay1(t *testing.T) {
	fmt.Println("===== Day 1 =====")

	{
		expected := Result{24000, 45000}
		actual := Run("example.txt")
		if actual != expected {
			t.Errorf("Expected: %d but got: %d\n", expected, actual)
		}
	}

	{
		expected := Result{69795, 208437}
		actual := Run("input.txt")
		if actual != expected {
			t.Errorf("Expected: %d but got: %d\n", expected, actual)
		}
	}
}
