package main

import "testing"

func runSample(t *testing.T, ts uint64, expect uint64) {
	res := solve(ts)

	if res != expect {
		t.Errorf("Sample %d, expect %d, but got %d", ts, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 0)
}

func TestSample2(t *testing.T) {
	runSample(t, 8, 0)
}

func TestSample3(t *testing.T) {
	runSample(t, 11, 5)
}
