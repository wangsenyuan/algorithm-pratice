package main

import "testing"

func runSample(t *testing.T, x, y int, expect int64) {
	res := solve(x, y)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 12345, 6789, 53384)
}
