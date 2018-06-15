package main

import "testing"

func runSample(t *testing.T, n, k, p int, A [][]int, P [][]int, expect bool) {
	res := solve(n, k, p, A, P)

	if res != expect {
		t.Errorf("Sample %d %d %d %v %v, expect %t, but got %t", n, k, p, A, P, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k, p := 10, 3, 2
	A := [][]int{
		{4, 5},
		{2, 3},
		{6, 1},
	}

	P := [][]int{
		{4, 5},
		{6, 3},
	}

	expect := true

	runSample(t, n, k, p, A, P, expect)
}

func TestSample2(t *testing.T) {
	n, k, p := 10, 3, 1
	A := [][]int{
		{2, 5},
		{10, 1},
		{6, 9},
	}
	P := [][]int{
		{1, 10},
	}
	expect := false

	runSample(t, n, k, p, A, P, expect)
}
