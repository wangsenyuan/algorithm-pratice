package main

import "testing"

func runSample(t *testing.T, low, high int64, expect int64) {
	res := solve(low, high)

	if res != expect {
		t.Errorf("Sample %d %d, expect %d, but got %d", low, high, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 11, 12, 14)
}

func TestSample2(t *testing.T) {
	runSample(t, 932451, 935212, 101245)
}
