package p939

import "testing"

func runSample(t *testing.T, points [][]int, expect int) {
	res := minAreaRect(points)
	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", points, expect, res)
	}
}

func TestSample1(t *testing.T) {
	points := [][]int{
		{1, 1}, {1, 3}, {3, 1}, {3, 3}, {2, 2},
	}
	runSample(t, points, 4)
}

func TestSample2(t *testing.T) {
	points := [][]int{
		{1, 1}, {1, 3}, {3, 1}, {3, 3}, {4, 1}, {4, 3},
	}
	runSample(t, points, 2)
}

func TestSample3(t *testing.T) {
	points := [][]int{
		{3, 2}, {3, 1}, {4, 4}, {1, 1}, {4, 3}, {0, 3}, {0, 2}, {4, 0},
	}
	runSample(t, points, 0)
}
