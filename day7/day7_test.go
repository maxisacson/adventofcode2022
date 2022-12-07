package day7

import (
	"fmt"
	"testing"
)

func TestDay7(t *testing.T) {
	expected := 95437
	actual := Run("example.txt")
	if actual != expected {
		t.Errorf("Expected: %d\nbut got: %d\n", expected, actual)
	}

	answer := 919137
	sum := Run("input.txt")
	if sum != answer {
		t.Errorf("Expected: %d\nbut got: %d\n", answer, sum)
	}
	fmt.Println("--- Part 1 ---")
	fmt.Println("Sum:", sum)

}
