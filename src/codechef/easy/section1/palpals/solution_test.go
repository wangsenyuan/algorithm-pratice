package main

import "testing"

func runSample(t *testing.T, s string, expect bool) {
	res := solve(s)

	if res != expect {
		t.Errorf("Sample %s, epxect %t, but got %t", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "acbbbadzdz", true)
}

func TestSample2(t *testing.T) {
	runSample(t, "abcd", false)
}

func TestSample3(t *testing.T) {
	runSample(t, "xyxyxy", true)
}

func TestSample4(t *testing.T) {
	runSample(t, "ghxxxyyzz", true)
}
