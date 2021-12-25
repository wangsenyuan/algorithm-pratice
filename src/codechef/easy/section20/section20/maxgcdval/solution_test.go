package main

import "testing"

func runSample(t *testing.T, k int, expect int) {
	res := int(solve(k))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 6)
}
