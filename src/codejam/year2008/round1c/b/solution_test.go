package main

import "testing"

func runSample(t *testing.T, s string, expect int64) {
	res := solve(s)
	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "1", 0)
	runSample(t, "9", 1)
	runSample(t, "011", 6)
	runSample(t, "12345", 64)
}
