package p2008

import "testing"

func runSample(t *testing.T, n int, rides [][]int, expect int64) {
	res := maxTaxiEarnings(n, rides)

	if expect != res {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	rides := [][]int{
		{2, 5, 4}, {1, 5, 1},
	}
	var expect int64 = 7
	runSample(t, n, rides, expect)
}

func TestSample2(t *testing.T) {
	n := 20
	rides := [][]int{
		{1, 6, 1}, {3, 10, 2}, {10, 12, 3}, {11, 12, 2}, {12, 15, 2}, {13, 18, 1},
	}
	var expect int64 = 20
	runSample(t, n, rides, expect)
}

func TestSample3(t *testing.T) {
	n := 10
	rides := [][]int{
		{9, 10, 2}, {4, 5, 6}, {6, 8, 1}, {1, 5, 5}, {4, 9, 5}, {1, 6, 5}, {4, 8, 3}, {4, 7, 10}, {1, 9, 8}, {2, 3, 5},
	}
	var expect int64 = 22
	runSample(t, n, rides, expect)
}
