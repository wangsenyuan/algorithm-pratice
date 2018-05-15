package main

import "testing"

func runSample(t *testing.T, N int64, p, q int64) {
	a, b := solve(N)
	if a != p || b != q {
		t.Errorf("sample %d, expect %d %d, but got %d %d", N, p, q, a, b)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 1, 10)
}

func TestSample2(t *testing.T) {
	runSample(t, 5, 1, 10)
}

func TestSample3(t *testing.T) {
	runSample(t, 11, 6, 55)
}
