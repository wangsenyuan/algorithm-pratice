package main

import "testing"

func runSample(t *testing.T, L, R int64, expect int) {
	res := solve(L, R)
	if res != expect {
		t.Errorf("sample: %d - %d, should give %d, but got %d", L, R, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 5, 13, 8)
}

func TestSample2(t *testing.T) {
	runSample(t, 1, 100, 90)
}
