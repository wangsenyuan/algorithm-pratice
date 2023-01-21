package main

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := solve(s)
	if res != expect {
		t.Errorf("Sample %s, expect %s, but got %s", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "R23C55", "BC23")
}
func TestSample2(t *testing.T) {
	runSample(t,  "BC23","R23C55")
}