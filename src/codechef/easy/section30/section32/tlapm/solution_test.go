package main

import "testing"

func runSample(t *testing.T, x1, y1, x2, y2 int, expect int) {
	res := solve(x1, y1, x2, y2)

	if res != int64(expect) {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 1, 3, 3, 32)
}
