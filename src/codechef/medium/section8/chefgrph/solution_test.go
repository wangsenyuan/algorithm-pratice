package main

import "testing"

func runSample(t *testing.T, n int64, m int, edges [][]int64, expect int) {
	res := solve(n, m, edges)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", n, m, edges, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 4, 2
	edges := [][]int64{
		{2, 1, 5, 0},
		{0, 0, 4, 0},
	}
	expect := 19
	runSample(t, int64(n), m, edges, expect)
}
