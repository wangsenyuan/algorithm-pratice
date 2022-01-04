package main

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := int(solve(n))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 0)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 10)
}

func TestSample3(t *testing.T) {
	runSample(t, 5, 270)
}
