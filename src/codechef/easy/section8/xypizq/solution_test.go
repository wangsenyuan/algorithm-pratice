package main

import "testing"

func runSample(t *testing.T, n int, tt int, x, y, z int, a, b int64) {
	c, d := solve(n, tt, x, y, z)
	if a != c || b != d {
		t.Errorf("Sample %d %d %d %d %d, expect %d %d, but got %d %d", n, tt, x, y, z, a, b, c, d)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 175456798, 1, 151163203, 151163204, 151163205, 66583464, 116971199)
}
