package day22

import (
	"fmt"
	"testing"
)

func TestDay22(t *testing.T) {
	fmt.Println("===== Day 22 =====")

	{
		expected := Result{6032, 5031}
		actual := Run("example.txt", 4)
		if actual != expected {
			t.Errorf("Expected: %d but got: %d\n", expected, actual)
		}
	}

	{
		expected := Result{181128, 52311}
		actual := Run("input.txt", 50)
		if actual != expected {
			t.Errorf("Expected: %d but got: %d\n", expected, actual)
		}
	}
}
