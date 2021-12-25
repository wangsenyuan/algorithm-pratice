package main

import "testing"

func runSample(t *testing.T, n int, X [][]int, expect int64) {
	res := solve(n, X)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	X := [][]int{
		{-2, 1},
		{2},
	}
	var expect int64 = 2

	runSample(t, n, X, expect)
}
