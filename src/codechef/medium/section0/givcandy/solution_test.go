package main

import "testing"

func runSample(t *testing.T, A, B, C, D, expect int) {
	res := solve(A, B, C, D)
	if res != expect {
		t.Errorf("%d, %d, %d, %d expect %d, but got %d", A, B, C, D, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 5, 10, 3, 0)
}

func TestSample2(t *testing.T) {
	runSample(t, 1, 2, 2, 2, 1)
}
