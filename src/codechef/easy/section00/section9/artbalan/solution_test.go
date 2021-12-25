package main

import "testing"

func runSample(t *testing.T, S string, expect int) {
	res := solve(S)
	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", S, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "ABCB", 1)
}

func TestSample2(t *testing.T) {
	runSample(t, "BBC", 1)
}

func TestSample3(t *testing.T) {
	runSample(t, "AAA", 0)
}
