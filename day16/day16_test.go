package day16

import (
	"fmt"
	"testing"
)

func TestDay16(t *testing.T) {
	fmt.Println("===== Day 16 =====")

	{
		expected := Result{1651, 0}
		actual := Run("example.txt")
		if actual != expected {
			t.Errorf("Expected: %d but got: %d\n", expected, actual)
		}
	}

	// {
	// 	expected := Result{1880, 0}
	// 	actual := Run("input.txt")
	// 	if actual != expected {
	// 		t.Errorf("Expected: %d but got: %d\n", expected, actual)
	// 	}
	// }
}
