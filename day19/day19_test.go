package day19

import (
	"fmt"
	"testing"
)

func TestDay19(t *testing.T) {
	fmt.Println("===== Day 19 =====")

	{
		expected := Result{9, 0}
		actual := Run("example.txt")
		if actual != expected {
			t.Errorf("Expected: %d but got: %d\n", expected, actual)
		}
	}

	// {
	// 	expected := Result{0, 0}
	// 	actual := Run("input.txt")
	// 	if actual != expected {
	// 		t.Errorf("Expected: %d but got: %d\n", expected, actual)
	// 	}
	// }
}
