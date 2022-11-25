package main

import "testing"

func runSample(t *testing.T, n, m int, expect int) {
	res := int(solve(n, m))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 5, 5)
}
