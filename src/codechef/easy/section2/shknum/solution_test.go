package main

import "testing"

func runSample(t *testing.T, N int, expect int) {
	res := solve(N)
	if res != expect {
		t.Errorf("Sample %d, expect %d, but got %d", N, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 0)
}

func TestSample2(t *testing.T) {
	runSample(t, 10, 0)
}

func TestSample3(t *testing.T) {
	runSample(t, 22, 2)
}

func TestSample4(t *testing.T) {
	runSample(t, 4, 1)
}
