package day20

import (
	"fmt"
	"testing"
)

func TestDay20(t *testing.T) {
	fmt.Println("===== Day 20 =====")

	{
		expected := Result{3, 1623178306}
		actual := Run("example.txt")
		if actual != expected {
			t.Errorf("Expected: %d but got: %d\n", expected, actual)
		}
	}

	{
		expected := Result{7004, 17200008919529}
		actual := Run("input.txt")
		if actual != expected {
			t.Errorf("Expected: %d but got: %d\n", expected, actual)
		}
	}
}
