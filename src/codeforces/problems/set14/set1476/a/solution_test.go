package main

import "testing"

func runSample(t *testing.T, n, k int, expect int) {
	res := solve(n, k)

	if res != expect {
		t.Errorf("Sample %d %d, expect %d, but got %d", n, k, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 5, 5)
}

func TestSample2(t *testing.T) {
	runSample(t, 4, 3, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, 8, 8, 1)
}

func TestSample4(t *testing.T) {
	runSample(t, 8, 17, 3)
}
