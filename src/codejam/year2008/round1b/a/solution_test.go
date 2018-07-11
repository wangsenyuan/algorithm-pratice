package main

import "testing"

func runSample(t *testing.T, n, A, B, C, D, x0, y0, M int, expect int) {
	res := solve(n, A, B, C, D, x0, y0, M)
	if res != expect {
		t.Errorf("Sample %d %d %d %d %d %d %d %d, expect %d, but got %d", n, A, B, C, D, x0, y0, M, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, A, B, C, D, x0, y0, M := 4, 10, 7, 1, 2, 0, 1, 20
	expect := 1
	runSample(t, n, A, B, C, D, x0, y0, M, expect)
}

func TestSample2(t *testing.T) {
	n, A, B, C, D, x0, y0, M := 6, 2, 0, 2, 1, 1, 2, 11
	expect := 2
	runSample(t, n, A, B, C, D, x0, y0, M, expect)
}
