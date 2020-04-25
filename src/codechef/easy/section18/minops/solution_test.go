package main

import "testing"

func runSample(t *testing.T, S, R string, expect int) {
	res := solve(S, R)

	if res != expect {
		t.Errorf("Sample %s %s, expect %d, but got %d", S, R, expect, res)
	}
}

func TestSample1(t *testing.T) {
	S := "adefb"
	R := "bdefa"
	runSample(t, S, R, 4)
}
