package main

import "testing"

func runSample(t *testing.T, l, r int64, expect int) {
	res := solve(l, r)
	if res != expect{
		t.Errorf("Sample %d %d, expect %d, but got %d", l, r, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 5, 27)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 4, 17)
}

func TestSample3(t *testing.T) {
	runSample(t, 100, 100, 441)
}
