package main

import "testing"

func runSample(t *testing.T, n int, kingdoms [][]int, expect int) {
	res := solve(n, kingdoms)
	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, kingdoms, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	kingdoms := [][]int{
		{1, 3},
		{2, 5},
		{6, 9},
	}
	expect := 2
	runSample(t, n, kingdoms, expect)
}
