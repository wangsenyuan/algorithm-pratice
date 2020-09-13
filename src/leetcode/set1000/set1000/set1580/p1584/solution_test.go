package p1584

import "testing"

func runSample(t *testing.T, points [][]int, expect int) {
	res := minCostConnectPoints(points)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	points := [][]int{
		{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0},
	}
	expect := 20
	runSample(t, points, expect)
}

func TestSample2(t *testing.T) {
	points := [][]int{
		{3, 12}, {-2, 5}, {-4, 1},
	}
	expect := 18
	runSample(t, points, expect)
}

func TestSample3(t *testing.T) {
	points := [][]int{
		{0, 0}, {1, 1}, {1, 0}, {-1, 1},
	}
	expect := 4
	runSample(t, points, expect)
}

func TestSample4(t *testing.T) {
	points := [][]int{
		{-1000000, -1000000}, {1000000, 1000000},
	}
	expect := 4000000
	runSample(t, points, expect)
}
