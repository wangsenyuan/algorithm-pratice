package main

import "testing"

func runSample(t *testing.T, N, M, X, K int, S string, expect bool) {
	res := solve(N, M, X, K, S)
	if res != expect {
		t.Errorf("Sample %d %d %d %d %v, expect %t, but got %t", N, M, X, K, S, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, M, X, K := 4, 4, 2, 4
	S := "EEOO"
	runSample(t, N, M, X, K, S, true)
}

func TestSample2(t *testing.T) {
	N, M, X, K := 4, 3, 1, 4
	S := "EEOO"
	runSample(t, N, M, X, K, S, false)
}
