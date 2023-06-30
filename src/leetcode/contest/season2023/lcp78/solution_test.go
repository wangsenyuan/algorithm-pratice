package lcp78

import "testing"

func runSample(t *testing.T, rampart [][]int, expect int) {
	res := rampartDefensiveLine(rampart)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	rampart := [][]int{
		{1, 2}, {5, 8}, {11, 15}, {18, 25},
	}
	expect := 4
	runSample(t, rampart, expect)
}
