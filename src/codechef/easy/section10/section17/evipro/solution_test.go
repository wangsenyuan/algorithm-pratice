package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := int(solve(s))

	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "001", 6)
}
