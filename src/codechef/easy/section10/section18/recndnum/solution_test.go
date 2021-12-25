package main

import "testing"

func runSample(t *testing.T, n, k int, expect int) {
	res := solve(n, k)

	if res != expect {
		t.Errorf("Sample %d %d, expect %d, but got %d", n, k, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 0, 1, 0)
}

func TestSample2(t *testing.T) {
	runSample(t, 1, 1, 1)
}

func TestSample3(t *testing.T) {
	runSample(t, 2, 1, 4)
}

func TestSample4(t *testing.T) {
	runSample(t, 1, 3, 5)
}

func TestSample5(t *testing.T) {
	runSample(t, 4, 6, 46)
}
