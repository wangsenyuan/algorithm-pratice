package main

import "testing"

func runSample(t *testing.T, A []int64, K int64, expect int64) {
	res := solve(A, K)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	var K int64 = 8
	A := []int64{1, 6, 4}
	var expect int64 = 2
	runSample(t, A, K, expect)
}
