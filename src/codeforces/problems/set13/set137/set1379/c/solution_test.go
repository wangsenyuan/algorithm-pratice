package main

import "testing"

func runSample(t *testing.T, n, m int, F [][]int, expect int64) {
	res := solve(n, m, F)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", n, m, F, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 4, 3
	F := [][]int{
		{5, 0},
		{1, 4},
		{2, 2},
	}
	var expect int64 = 14

	runSample(t, n, m, F, expect)
}

func TestSample2(t *testing.T) {
	n, m := 5, 3
	F := [][]int{
		{5, 2},
		{4, 2},
		{3, 1},
	}
	var expect int64 = 16

	runSample(t, n, m, F, expect)
}
