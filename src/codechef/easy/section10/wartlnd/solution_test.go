package main

import "testing"

func runSample(t *testing.T, n int, roads [][]int, expect int64) {
	res := solve(n, roads)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, roads, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	roads := [][]int{
		{1, 2, 4},
		{1, 3, 3},
		{2, 4, 1},
		{2, 5, 2},
	}
	expect := int64(17)
	runSample(t, n, roads, expect)
}
