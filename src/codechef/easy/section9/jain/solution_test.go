package main

import "testing"

func runSample(t *testing.T, n int, ss []string, expect int64) {
	res := solve(n, ss)
	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, ss, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	ss := []string{
		"aaooaoaooa",
		"uiieieiieieuuu",
		"aeioooeeiiaiei",
	}
	runSample(t, n, ss, 2)
}
