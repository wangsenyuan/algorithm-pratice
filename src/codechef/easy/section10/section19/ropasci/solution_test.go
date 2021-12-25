package main

import "testing"

func runSample(t *testing.T, S string, expect string) {
	res := solve(len(S), S)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	S := "SSPR"
	W := "RRPR"
	runSample(t, S, W)
}
