package main

import "testing"

func runSample(t *testing.T, S string, X byte, expect int64) {
	res := solve(S, X)
	if res != expect {
		t.Errorf("Sample %s %b, expect %d, but got %d", S, X, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "abb", 'b', 5)
}

func TestSample2(t *testing.T) {
	runSample(t, "abcabc", 'c', 15)
}
