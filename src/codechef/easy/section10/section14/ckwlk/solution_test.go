package main

import "testing"

func runSample(t *testing.T, s string, expect bool) {
	res := solve(s)
	if res != expect {
		t.Errorf("Sample %s, expect %t, but got %t", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "200", true)
}

func TestSample2(t *testing.T) {
	runSample(t, "90", false)
}

func TestSample3(t *testing.T) {
	runSample(t, "1000000000000", true)
}

func TestSample4(t *testing.T) {
	runSample(t, "1024000000000", false)
}
