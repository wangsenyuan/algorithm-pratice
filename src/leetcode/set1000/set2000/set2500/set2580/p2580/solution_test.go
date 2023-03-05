package p2580

import "testing"

func runSample(t *testing.T, ranges [][]int, expect int) {
	res := countWays(ranges)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	ranges := [][]int{
		{1, 3}, {10, 20}, {2, 5}, {4, 8},
	}
	expect := 4
	runSample(t, ranges, expect)
}

func TestSample2(t *testing.T) {
	ranges := [][]int{
		{5, 11}, {20, 22}, {1, 3}, {21, 22}, {11, 11},
	}
	expect := 8
	runSample(t, ranges, expect)
}
