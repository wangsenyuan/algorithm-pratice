package main

import "testing"

func runSample(t *testing.T, x, y int64, expect int) {
	res := solve(x, y)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 7, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, 1, 4, 4)
}
