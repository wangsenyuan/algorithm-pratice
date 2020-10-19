package main

import "testing"

func runSample(t *testing.T, n, k, x int, expect bool) {
	res := solve(n, k, x)

	if res != expect {
		t.Errorf("Sample %d %d %d, expect %t, but got %t", n, k, x, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 4, 3, 3, true)
}

func TestSample2(t *testing.T) {
	runSample(t, 5, 8, 1, false)
}

func TestSample3(t *testing.T) {
	runSample(t, 10, 8, 2, true)
}

func TestSample4(t *testing.T) {
	runSample(t, 19, 100, 34, false)
}

func TestSample5(t *testing.T) {
	runSample(t, 69, 100, 45, true)
}
