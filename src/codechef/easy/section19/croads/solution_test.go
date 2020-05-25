package main

import "testing"

func runSample(t *testing.T, N int, expect int64) {
	res := solve(N)

	if res != expect {
		t.Errorf("Sample %d expect %d, but got %d", N, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, 5, 8)
}

func TestSample3(t *testing.T) {
	runSample(t, 11, 23)
}

func TestSample4(t *testing.T) {
	runSample(t, 8, -1)
}
