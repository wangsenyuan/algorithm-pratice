package main

import "testing"

func runSample(t *testing.T, n, m int, expect int64) {
	res := solve(n, m)

	if res != expect {
		t.Errorf("Sample %d %d, expect %d, but got %d", n, m, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 1, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 1, 3, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, 2, 2, 2)
}

func TestSample4(t *testing.T) {
	runSample(t, 3, 3, 5)
}

func TestSample5(t *testing.T) {
	runSample(t, 5, 3, 8)
}
