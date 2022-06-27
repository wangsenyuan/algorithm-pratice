package p2316

import "testing"

func runSample(t *testing.T, n int, edges [][]int, expect int64) {
	res := countPairs(n, edges)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 7
	edges := [][]int{
		{0, 2}, {0, 5}, {2, 4}, {1, 6}, {5, 4},
	}
	var expect int64 = 14
	runSample(t, n, edges, expect)
}
