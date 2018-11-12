package main

import "testing"

func runSample(t *testing.T, n int, x, y int) {
	u, v := solve(n)

	if u != x || v != y {
		t.Errorf("Sample %d, expect %d/%d, but got %d/%d", n, x, y, u, v)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 1, 2)
}
func TestSample2(t *testing.T) {
	runSample(t, 2, 1, 4)
}
