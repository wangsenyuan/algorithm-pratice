package main

import "testing"

func runSample(t *testing.T, x, y uint64, expect int64) {
	can, res := solve(x, y)
	if can && res != uint64(expect) {
		t.Errorf("Sample %d %d, expect %d, but got %d", x, y, expect, res)
	} else if !can && expect > 0 {
		t.Errorf("Sample %d %d, expect %d, but got %d", x, y, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 3, 21)
}

func TestSample2(t *testing.T) {
	runSample(t, 12, 5, 120)
}
