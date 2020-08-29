package main

import "testing"

func runSample(t *testing.T, n, k int, expect int) {
	res := solve(n, k)

	if res != expect {
		t.Errorf("Sample %d %dm expect %d, but got %d", n, k, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 7, 3, 16)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 7, 0)
}

func TestSample3(t *testing.T) {
	runSample(t, 1337, 42, 95147305)
}

func TestSample4(t *testing.T) {
	runSample(t, 5000, 1, 5000)
}
