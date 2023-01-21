package main

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := solve(n)

	if res != expect {
		t.Errorf("Sample %d, expect %d, but got %d", n, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 4)
}

func TestSample2(t *testing.T) {
	runSample(t, 4, 4)
}

func TestSample3(t *testing.T) {
	runSample(t, 5, 12)
}

func TestSample4(t *testing.T) {
	runSample(t, 100, 990998587)
}

func TestSample5(t *testing.T) {
	runSample(t, 2000000, 804665184)
}
