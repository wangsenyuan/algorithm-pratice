package main

import "testing"

func runSample(t *testing.T, n, m int, A [][]int, expect int) {
	res := solve(n, m, A)
	if expect != res {
		t.Errorf("sample: %v, should give %d, but got %d", A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 2, 2
	A := [][]int{
		{1, 1},
		{1, 0},
	}

	runSample(t, n, m, A, 3)
}

func TestSample2(t *testing.T) {
	n, m := 3, 4
	A := [][]int{
		{3, 4, 7, 8},
		{4, 2, 6, 8},
		{7, 1, 7, 3},
	}
	runSample(t, n, m, A, 5)
}
