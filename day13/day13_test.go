package day13

import (
	"fmt"
	"testing"
)

func TestDay13(t *testing.T) {
	fmt.Println("===== Day 13 =====")

	{
		expected := Result{13, 140}
		actual := Run("example.txt")
		if actual != expected {
			t.Errorf("Expected: %d but got: %d\n", expected, actual)
		}
	}

	{
		expected := Result{4809, 22600}
		actual := Run("input.txt")
		if actual != expected {
			t.Errorf("Expected: %d but got: %d\n", expected, actual)
		}
	}
}
