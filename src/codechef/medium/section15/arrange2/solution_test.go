package main

import "testing"

func runSample(t *testing.T, A, B int64, expect int) {
	res := solve(A, B)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 499875921, 499875921, 1)
}
