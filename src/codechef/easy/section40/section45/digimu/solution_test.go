package main

import "testing"

func runSample(t *testing.T, A, B, K int64, expect int64) {
	res := solve(A, B, K)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	var A, B, K int64 = 10, 100, 4
	var expect int64 = 59
	runSample(t, A, B, K, expect)
}

func TestSample2(t *testing.T) {
	var A, B, K int64 = 10, 58, 4
	var expect int64 = -1
	runSample(t, A, B, K, expect)
}
