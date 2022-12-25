package day25

import (
	"fmt"
	"testing"
)

func TestDay25(t *testing.T) {
	fmt.Println("===== Day 25 =====")

	{
		expected := Result{"2=-1=0", ""}
		actual := Run("example.txt")
		if actual != expected {
			t.Errorf("Expected: %s but got: %s\n", expected, actual)
		}
	}

	{
		expected := Result{"2-0==21--=0==2201==2", ""}
		actual := Run("input.txt")
		if actual != expected {
			t.Errorf("Expected: %s but got: %s\n", expected, actual)
		}
	}
}
