package main

import "testing"

func runSample(t *testing.T, n, m int, segs [][]int, expect int) {
	res := solve(n, m, segs)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", n, m, segs, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 3, 3
	segs := [][]int{
		{1, 1},
		{2, 2},
		{3, 3},
	}
	expect := 8
	runSample(t, n, m, segs, expect)
}
