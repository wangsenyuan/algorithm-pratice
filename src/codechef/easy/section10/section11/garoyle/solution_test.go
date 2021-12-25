package main

import "testing"

func runSample(t *testing.T, S []string, expect int) {
	res := solve(S)
	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", S, expect, res)
	}
}


func TestSample1(t *testing.T) {
	S := []string{
		"TTFFF",
		"TTFFF",
		"TTTFF",
		"FFFTT",
		"FFFTT",
	}

	runSample(t, S, 2)
}