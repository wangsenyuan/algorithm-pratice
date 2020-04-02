package main

import "testing"

func runSample(t *testing.T, s string, A int, K int, expect int) {
	res := solve(s, A, K)

	if res != expect {
		t.Errorf("Sample %s %d %d, expect %d, but got %d", s, A, K, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "aba", 2, 2, 5)
}
