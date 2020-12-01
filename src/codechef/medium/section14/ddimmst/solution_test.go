package main

import "testing"

func runSample(t *testing.T, n int, d int, X [][]int, expect int64) {
	res := solve(n, d, X)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, d := 2, 2
	X := [][]int{
		{1, 1},
		{2, 2},
	}
	var expect int64 = 2
	runSample(t, n, d, X, expect)
}
