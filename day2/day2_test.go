package day2

import (
	"fmt"
	"testing"
)

func TestDay2(t *testing.T) {
	fmt.Println("===== Day 2 =====")

	{
		expected := Result{15, 12}
		actual := Run("example.txt")
		if actual != expected {
			t.Errorf("Expected: %d but got: %d\n", expected, actual)
		}
	}

	{
		expected := Result{11666, 12767}
		actual := Run("input.txt")
		if actual != expected {
			t.Errorf("Expected: %d but got: %d\n", expected, actual)
		}
	}
}
