package main

import "testing"

func runSample(t *testing.T, n, m, q int, operations [][]int, expect int64) {
	res := solve(n, m, q, operations)

	if res != expect {
		t.Errorf("Sample %d %d %d %v, expect %d, but got %d", n, m, q, operations, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m, q := 2, 2, 3
	ops := [][]int{
		{1, 1},
		{1, 2},
		{2, 1},
	}
	runSample(t, n, m, q, ops, 2)
}
