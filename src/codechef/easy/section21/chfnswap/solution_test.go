package main

import "testing"

func runSample(t *testing.T, n int, expect int64) {
	res := solve(n)

	if res != expect {
		t.Errorf("Sample %d, expect %d, but got %d", n, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 4, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, 7, 3)
}
