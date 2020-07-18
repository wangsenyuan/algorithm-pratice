package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, expect int64) {
	res := solve(n, E)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, E, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	E := [][]int{
		{1, 2},
		{3, 5},
		{3, 4},
		{2, 3},
		{5, 6},
	}
	var expect int64 = 8

	runSample(t, n, E, expect)
}
