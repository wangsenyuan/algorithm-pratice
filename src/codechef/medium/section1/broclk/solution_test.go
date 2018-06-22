package main

import "testing"

func runSample(xt *testing.T, l, d int, t uint64, expect int64) {
	res := solve(l, d, t)

	if res != expect {
		xt.Errorf("Sample %d %d %d, expect %d, but got %d", l, d, t, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 4, 2, 1, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 4, 2, 2, 1000000005)
}

func TestSample3(t *testing.T) {
	runSample(t, 4, 2, 3, 1000000003)
}
