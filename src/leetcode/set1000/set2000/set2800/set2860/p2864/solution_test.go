package p2864

import "testing"

func runSample(t *testing.T, n int, edges [][]int, expect int64) {
	res := countPaths(n, edges)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	edges := [][]int{
		{1, 2}, {1, 3}, {2, 4}, {2, 5},
	}
	var expect int64 = 4
	runSample(t, n, edges, expect)
}

func TestSample2(t *testing.T) {
	n := 6
	edges := [][]int{
		{1, 2}, {1, 3}, {2, 4}, {3, 5}, {3, 6},
	}
	var expect int64 = 6
	runSample(t, n, edges, expect)
}

func TestSample3(t *testing.T) {
	n := 6
	edges := [][]int{
		{4, 1}, {5, 4}, {2, 1}, {3, 4},
	}
	var expect int64 = 6
	runSample(t, n, edges, expect)
}
