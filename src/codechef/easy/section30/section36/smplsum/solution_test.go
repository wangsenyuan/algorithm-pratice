package main

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := int(solve(n))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 1)
	runSample(t, 2, 3)
	runSample(t, 3, 7)
	runSample(t, 4, 11)
	runSample(t, 5, 21)
}
