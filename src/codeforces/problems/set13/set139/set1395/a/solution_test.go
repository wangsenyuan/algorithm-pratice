package main

import "testing"

func runSample(t *testing.T, r, g, b, w int, expect bool) {
	res := solve(r, g, b, w)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 0, 1, 1, 1, false)
}

func TestSample2(t *testing.T) {
	runSample(t, 8, 1, 9, 3, true)
}

func TestSample3(t *testing.T) {
	runSample(t, 1000000000, 1000000000, 1000000000, 1000000000, true)
}
