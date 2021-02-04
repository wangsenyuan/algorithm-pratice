package main

import "testing"

func runSample(t *testing.T, n int, k int, E [][]int, expect int64) {
	res := solve(n, k, E)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 3, 1
	E := [][]int{
		{0, 1},
		{1, 2},
	}
	var expect int64 = 5
	runSample(t, n, k, E, expect)
}

func TestSample2(t *testing.T) {
	n, k := 6, 3
	E := [][]int{
		{0, 1},
		{1, 2},
		{2, 3},
		{2, 4},
		{3, 5},
	}
	var expect int64 = 23
	runSample(t, n, k, E, expect)
}
