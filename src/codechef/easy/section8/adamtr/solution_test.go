package main

import "testing"

func runSample(t *testing.T, n int, A [][]int, B [][]int, expect bool) {
	res := solve(n, A, B)
	if res != expect {
		t.Errorf("Sample %d %v %v, expect %t, but got %t", n, A, B, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	A := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	B := [][]int{
		{1, 2, 7},
		{4, 5, 8},
		{3, 6, 9},
	}
	runSample(t, n, A, B, true)
}
