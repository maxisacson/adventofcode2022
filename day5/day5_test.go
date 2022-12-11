package day5

import (
	"fmt"
	"testing"
)

func TestDay5(t *testing.T) {
	fmt.Println("===== Day 5 =====")

	{
		expected := Result{"CMZ", "MCD"}
		actual := Run("example.txt")
		if actual != expected {
			t.Errorf("Expected: %s but got: %s\n", expected, actual)
		}
	}

	{
		expected := Result{"VRWBSFZWM", "RBTWJWMCF"}
		actual := Run("input.txt")
		if actual != expected {
			t.Errorf("Expected: %s but got: %s\n", expected, actual)
		}
	}
}
