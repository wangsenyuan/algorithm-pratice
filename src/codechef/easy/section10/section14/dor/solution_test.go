package main

import "testing"

func runSample(t *testing.T, L, R uint64, expect uint64) {
	res := solve(L, R)
	if res != expect {
		t.Errorf("Sample %d %d, expect %d, but got %d", L, R, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1000000000123, 1000000123456, 1000000192511)
}

func TestSample2(t *testing.T) {
	runSample(t, 1031, 93714, 131071)
}
