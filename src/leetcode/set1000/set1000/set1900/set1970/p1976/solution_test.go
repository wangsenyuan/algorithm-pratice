package p1976

import "testing"

func runSample(t *testing.T, n int, roads [][]int, expect int) {
	res := countPaths(n, roads)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 7
	roads := [][]int{
		{0, 6, 7}, {0, 1, 2}, {1, 2, 3}, {1, 3, 3}, {6, 3, 3}, {3, 5, 1}, {6, 5, 1}, {2, 5, 1}, {0, 4, 5}, {4, 6, 2},
	}
	expect := 4
	runSample(t, n, roads, expect)
}
