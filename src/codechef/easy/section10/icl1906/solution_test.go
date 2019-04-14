package main

import "testing"

func runSample(t *testing.T, D, H int, expect int) {
	res := solve(D, H)
	if res != expect {
		t.Errorf("Sample %d %d, expect %d, but got %d", D, H, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 9, 12)
}

func TestSample2(t *testing.T) {
	runSample(t, 13, 14, 15)
}
