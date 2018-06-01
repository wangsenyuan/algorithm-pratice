package main

import "testing"

func runSample(t *testing.T, n int, K int, A []int, expect int64) {
	res := solve(n, K, A)
	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", n, K, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, K := 4, 2
	A := []int{1, 2, 1, 2}
	runSample(t, n, K, A, 2)
}
