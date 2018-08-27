package main

import "testing"

func runSample(t *testing.T, N, K int, x, D int64, P []int64, expect int64) {
	res := solve(N, K, x, D, P)
	if res != expect {
		t.Errorf("Sample %d %d %d %d %v, expect %d, but got %d", N, K, x, D, P, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, K := 4, 3
	var x int64 = 5
	var D int64 = 3
	P := []int64{2, 1, 5}
	expect := int64(12)
	runSample(t, N, K, x, D, P, expect)
}

func TestSample2(t *testing.T) {
	N, K := 3, 2
	var x int64 = 8
	var D int64 = 2
	P := []int64{3, 8}
	expect := int64(-1)
	runSample(t, N, K, x, D, P, expect)
}
