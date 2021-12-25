package main

import "testing"

func runSample(t *testing.T, A, B int, expect int) {
	res := solve(A, B)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 11, 6)
}

func TestSample2(t *testing.T) {
	runSample(t, 13, 6, -3)
}
