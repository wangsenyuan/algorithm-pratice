package main

import "testing"

func runSample(t *testing.T, n int, A []int, M int, C [][]int, expect int) {
	res := solve(n, A, M, C)

	if res != expect {
		t.Errorf("Sample %d %v %d %v, expect %d, but got %d", n, A, M, C, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	A := []int{1, 2, 3, 4}
	M := 3
	C := [][]int{
		{2, 1, 2},
		{2, 3, 4},
		{3, 1, 2, 3},
	}
	expect := 6
	runSample(t, n, A, M, C, expect)
}
