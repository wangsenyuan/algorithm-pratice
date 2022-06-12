package main

import "testing"

func runSample(t *testing.T, n, k int, expect int) {
	res := int(solve(n, k))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 7, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 6, 4)
}
