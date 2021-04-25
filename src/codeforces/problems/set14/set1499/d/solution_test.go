package main

import "testing"

func runSample(t *testing.T, c, d, x, expect int) {
	res := solve(c, d, x)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 7, 25, 8)
}
