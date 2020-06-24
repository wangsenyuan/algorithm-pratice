package main

import "testing"

func runSample(t *testing.T, N, M int64, A, B int64) {
	C, D := solve(N, M)

	if C != A || D != B {
		t.Errorf("Sample expect %d %d, but got %d %d", A, B, C, D)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 100, 3, 76, 49152)
}
