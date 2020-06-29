package p1498

import "testing"

func runSample(t *testing.T, points [][]int, k int, expect int) {
	res := findMaxValueOfEquation(points, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	points := [][]int{
		{1, 3}, {2, 0}, {5, 10}, {6, -10},
	}
	k := 1
	expect := 4
	runSample(t, points, k, expect)
}
