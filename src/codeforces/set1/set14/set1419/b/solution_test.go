package main

import "testing"

func runSample(t *testing.T, x int64, expect int) {
	res := solve(x)

	if res != expect {
		t.Errorf("Sample %d, expect %d, but got %d", x, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 8, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, 6, 1)
}

func TestSample4(t *testing.T) {
	runSample(t, 1000000000000000000, 30)
}
