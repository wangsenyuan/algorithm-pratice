package main

import "testing"

func runSample(t *testing.T, n int, p, q int, A []int, x, y, z int) {
	u, v, w := solve(n, p, q, A)
	if u != x || v != y || w != z {
		t.Errorf("Sample %d %d %d %v, expect %d %d %d, but got %d %d %d", n, p, q, A, x, y, z, u, v, w)
	}
}

func TestSample1(t *testing.T) {
	n, p, q := 4, 1, 2
	A := []int{6, 6, 6, 7}
	x, y, z := 3, 1, 2
	runSample(t, n, p, q, A, x, y, z)
}

func TestSample2(t *testing.T) {
	n, p, q := 3, 5, 7
	A := []int{1, 3, 8}
	x, y, z := -1, -1, -1
	runSample(t, n, p, q, A, x, y, z)
}
