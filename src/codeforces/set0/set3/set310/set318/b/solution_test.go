package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)

	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "heavymetalisheavymetal"
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "trueheavymetalissotruewellitisalsosoheavythatyoucanalmostfeeltheweightofmetalonyou"
	expect := 3
	runSample(t, s, expect)
}
