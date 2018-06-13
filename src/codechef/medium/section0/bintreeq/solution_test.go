package main

import "testing"

func runSample(t *testing.T, n int, u int, v int, expect int) {
	res := solve(n, u, v)

	if res != expect {
		t.Errorf("Sample %d %d %d, expect %d, but got %d", n, u, v, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 11, 9, 11, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 10, 2, 2, 10)
}

func TestSample3(t *testing.T) {
	runSample(t, 8, 1, 8, 1)
}
