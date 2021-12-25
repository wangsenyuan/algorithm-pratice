package main

import "testing"

func runSample(t *testing.T, n int, u, v int, expect int) {
	res := solve(n, u, v)
	if res != expect {
		t.Errorf("Sample %d %d %d, expect %d, but got %d", n, u, v, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 20, 2, 3, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 20, 2, 6, 1)
}

func TestSample3(t *testing.T) {
	runSample(t, 20, 2, 9, 3)
}

func TestSample4(t *testing.T) {
	runSample(t, 25, 10, 25, 2)
}
