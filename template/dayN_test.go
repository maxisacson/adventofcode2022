package dayN

import (
	"fmt"
	"testing"
)

func TestDayN(t *testing.T) {
	fmt.Println("===== Day N =====")

	{
		expected := Result{0, 0}
		actual := Run("example.txt")
		if actual != expected {
			t.Errorf("Expected: %d but got: %d\n", expected, actual)
		}
	}

	{
		expected := Result{0, 0}
		actual := Run("input.txt")
		if actual != expected {
			t.Errorf("Expected: %d but got: %d\n", expected, actual)
		}
	}
}
