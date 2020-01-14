package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)

	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "3^40|10^2", 43)
}

func TestSample2(t *testing.T) {
	runSample(t, "92^95|56&2&3", 95)
}
