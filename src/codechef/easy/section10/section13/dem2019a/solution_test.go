package main

import "testing"

func runSample(tst *testing.T, n, m, t int, X [][]int, Y [][]int, base []int, expect int) {
	res := solve(n, m, t, X, Y, base)
	if res != expect {
		tst.Errorf("Sample %d %d %d %v %v %v, expect %d, but got %d", n, m, t, X, Y, base, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, M, T := 2, 2, 35
	X := [][]int{
		{1, 2, 15},
		{2, 10, 20},
	}
	Y := [][]int{
		{2, 1, 8},
		{5, 6, 12},
	}

	base := []int{5, 5}

	expect := 2

	runSample(t, N, M, T, X, Y, base, expect)
}
