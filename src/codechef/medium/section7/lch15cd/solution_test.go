package main

import "testing"

func runSample(t *testing.T, n, d int, expect int64) {
	res := solve(n, d)

	if res != expect {
		t.Errorf("Sample %d, %d, expect %d, but got %d", n, d, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 2, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 1, 3, 5)
}

func TestSample4(t *testing.T) {
	runSample(t, 2, 2, 1)
}

func TestSample5(t *testing.T) {
	runSample(t, 2, 3, 10)
}
