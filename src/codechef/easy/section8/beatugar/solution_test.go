package main

import "testing"

func runSample(t *testing.T, s string, expect bool) {
	res := solve(s)
	if res != expect {
		t.Errorf("Sample %s, expect %t, but got %t", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "RG", true)
}

func TestSample2(t *testing.T) {
	runSample(t, "RRGG", true)
}

func TestSample3(t *testing.T) {
	runSample(t, "RR", false)
}
