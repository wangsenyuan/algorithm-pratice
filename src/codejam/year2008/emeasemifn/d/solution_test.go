package main

import "testing"

func runSample(t *testing.T, n, k, p int, expect int) {
	res := solve(n, k, p)
	if res != expect {
		t.Errorf("Sample %d %d %d, expect %d, but got %d", n, k, p, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 10, 3, 3, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 5, 2, 3, 3)
}

func TestSample3(t *testing.T) {
	runSample(t, 40, 4, 8, 7380)
}
