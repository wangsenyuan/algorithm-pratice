package main

import "testing"

func runSample(t *testing.T, n, k int, A [][]int, expect int) {
	res := solve(n, k, A)
	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", n, k, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 3, 4
	A := [][]int{
		{1, 2, 3, 4},
		{2, 3, 4, 6},
		{6, 5, 4, 3},
	}
	expect := 2
	runSample(t, n, k, A, expect)
}

func TestSample2(t *testing.T) {
	n, k := 3, 3
	A := [][]int{
		{5, 5, 5},
		{4, 4, 6},
		{4, 5, 4},
	}
	expect := 3
	runSample(t, n, k, A, expect)
}

func TestSample3(t *testing.T) {
	n, k := 5, 2
	A := [][]int{
		{1, 1},
		{2, 2},
		{5, 4},
		{4, 4},
		{4, 1},
	}
	expect := 2
	runSample(t, n, k, A, expect)
}
