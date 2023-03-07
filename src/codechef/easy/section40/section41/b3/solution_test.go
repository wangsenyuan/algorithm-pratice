package main

import "testing"

func runSample(t *testing.T, y, x1, x2 int, expect int) {
	res := solve(x1, x2, y)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, -2, 1, 2)
}
