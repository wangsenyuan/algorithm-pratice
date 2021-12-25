package main

import "testing"

func runSample(t *testing.T, n int, S string, expect int) {
	res := solve(n, S)

	if res != expect {
		t.Errorf("Sample %d %s, expect %d, but got %d", n, S, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 4, "anxa", 1)
}
