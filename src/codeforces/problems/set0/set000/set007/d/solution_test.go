package main

import "testing"

func runSample(t *testing.T, s string, expect int64) {
	res := solve(s)

	if res != expect {
		t.Errorf("Sample %s, expect %d but go %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "abacaba", 6)
}

func TestSample2(t *testing.T) {
	runSample(t, "CCeCeCCCee", 4)
}
