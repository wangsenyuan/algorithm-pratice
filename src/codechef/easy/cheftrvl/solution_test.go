package main

import "testing"

func runSample(t *testing.T, n int, pairs [][]int, A []int, expect int64) {
	res := solve(n, pairs, A)
	if res != expect {
		t.Errorf("Sample %d %v %v, expect %d, but got %d", n, pairs, A, expect, res)
	}
}

func TestSample(t *testing.T) {
	n := 3
	pairs := [][]int{
		{1, 2},
		{1, 3},
	}
	A := []int{1, 2, 3}
	var expect int64 = 2
	runSample(t, n, pairs, A, expect)
}
