package main

import "testing"

func runSample(t *testing.T, n int, a int, k int, p int, q int) {
	x, y := solve(n, a, k)

	if int64(p) != x || int64(q) != y {
		t.Errorf("Sample expect %d, %d, but got %d %d", p, q, x, y)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 30, 2, 60, 1)
}
