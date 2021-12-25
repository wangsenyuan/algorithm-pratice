package main

import "testing"

func runSample(t *testing.T, n int, A [][]int, expect int) {
	res := solve(n, A)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	A := [][]int{
		{4, 5, 2, 3, 4},
		{2, 1, 6},
	}
	expect := 8
	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	A := [][]int{
		{4, 5, 2, 3, 4},
		{3, 1, 6, 7},
	}
	expect := 14
	runSample(t, n, A, expect)
}
