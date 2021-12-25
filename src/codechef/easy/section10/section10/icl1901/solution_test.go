package main

import "testing"

func runSample(t *testing.T, K, N int, expect int) {
	res := solve(K, N)
	if res != expect {
		t.Errorf("Sample %d %d, expect %d, but got %d", K, N, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 123, 5, 27)
}
