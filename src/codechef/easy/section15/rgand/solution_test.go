package main

import "testing"

func runSample(t *testing.T, l, r int64, expect int64) {
	res := solve(l, r)

	if res != expect {
		t.Errorf("Sample %d %d, expect %d, but got %d", l, r, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 4, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 4, 10, 16)
}
