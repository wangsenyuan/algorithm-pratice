package main

import "testing"

func runSample(t *testing.T, A int, N int, expect int) {
	res := solve(A, N)

	if res != expect {
		t.Errorf("Sample %d %d, expect %d, but got %d", A, N, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 3, 511620149)
}
