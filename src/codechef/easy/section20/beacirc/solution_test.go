package main

import "testing"

func runSample(t *testing.T, n int, circles [][]int, expect int64) {
	res := solve(n, circles)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", circles, expect, res)
	}
}

func TestSample1(t *testing.T) {
	cs := [][]int{
		{0, 0, 8},
		{0, 10, 6},
		{0, 5, 5},
		{0, 0, 8},
	}
	var expect int64 = 2
	runSample(t, len(cs), cs, expect)
}
