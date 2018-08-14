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
		{3, 4},
		{4, 5},
	}
	expect := 0
	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 1
	A := [][]int{
		{1, 2, 3},
	}
	expect := -1
	runSample(t, n, A, expect)
}
