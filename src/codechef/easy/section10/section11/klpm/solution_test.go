package main

import "testing"

func runSample(t *testing.T, S string, expect int64) {
	res := solve(S)
	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", S, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "abba", 7)
}
