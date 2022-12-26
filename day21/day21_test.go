package day21

import (
	"fmt"
	"testing"
)

func TestDay21(t *testing.T) {
	fmt.Println("===== Day 21 =====")

	{
		expected := Result{152, 301}
		actual := Run("example.txt")
		if actual != expected {
			t.Errorf("Expected: %d but got: %d\n", expected, actual)
		}
	}

	{
		expected := Result{93813115694560, 3910938071092}
		actual := Run("input.txt")
		if actual != expected {
			t.Errorf("Expected: %d but got: %d\n", expected, actual)
		}
	}
}
