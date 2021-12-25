package main

import "testing"

func runSample(t *testing.T, n, x int, P [][]int, a, b int) {
	c, d := solve(n, x, P)

	if c != a || b != d {
		t.Errorf("Sample expect %d %d, but got %d %d", a, b, c, d)
	}
}

func TestSample1(t *testing.T) {
	P := [][]int{
		{1, 1},
		{1, 0},
		{3, 2},
	}
	a, b := 2, 2
	runSample(t, 3, 1, P, a, b)
}
