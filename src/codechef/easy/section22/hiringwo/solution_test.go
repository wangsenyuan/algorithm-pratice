package main

import "testing"

func runSample(t *testing.T, k, x int, expect int) {
	res := solve(k, x)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 3, 4)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 6, 5)
}
