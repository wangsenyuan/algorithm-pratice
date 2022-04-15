package main

import "testing"

func runSample(t *testing.T, A, B uint64, expect int) {
	res := solve(A, B)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 6, 0)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 2, 1)
}
