package main

import "testing"

func runSample(t *testing.T, L, R int, expect int) {
	res := solve(L, R)
	if res != expect {
		t.Errorf("Sample %d %d, expect %d, but got %d", L, R, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 10, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, 11, 33, 8)
}
