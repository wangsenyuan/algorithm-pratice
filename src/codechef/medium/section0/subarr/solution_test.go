package main

import "testing"

func runSample(t *testing.T, n int, K int64, A []int64, expect int64) {
	res := solve(n, K, A)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", n, K, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	var K int64 = 2
	A := []int64{1, 2, 3, -1}
	var expect int64 = 4
	runSample(t, n, K, A, expect)
}
