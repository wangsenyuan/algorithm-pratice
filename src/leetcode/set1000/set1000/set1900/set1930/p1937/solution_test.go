package p1937

import "testing"

func runSample(t *testing.T, points [][]int, expect int64) {
	res := maxPoints(points)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	points := [][]int{
		{1, 2, 3}, {1, 5, 1}, {3, 1, 1},
	}
	var expect int64 = 9
	runSample(t, points, expect)
}

func TestSample2(t *testing.T) {
	points := [][]int{
		{1, 5}, {2, 3}, {4, 2},
	}
	var expect int64 = 11
	runSample(t, points, expect)
}

func TestSample3(t *testing.T) {
	points := [][]int{
		{0, 4, 3, 0, 0},
	}
	var expect int64 = 4
	runSample(t, points, expect)
}
