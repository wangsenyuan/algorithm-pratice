package main

import "testing"

func runSample(t *testing.T, k int, S []string, expect int) {
	res := solve(k, S)
	if len(res) != expect {
		t.Errorf("Sample expect %d length string, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 3
	S := []string{
		"BG",
		"GBB",
		"BGB",
	}
	expect := 4
	runSample(t, k, S, expect)
}
