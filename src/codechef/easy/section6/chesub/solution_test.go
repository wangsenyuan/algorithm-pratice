package main

import "testing"

func runSample(t *testing.T, N, K int, A []int, expect int64) {
	res := solve(N, K, A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, K := 5, 2
	A := []int{1, 2, -1, 3, 1}
	var expect int64 = 11
	runSample(t, N, K, A, expect)
}

func TestSample2(t *testing.T) {
	N, K := 5, 2
	A := []int{-1, 2, 11, -23, 12}
	var expect int64 = 37
	runSample(t, N, K, A, expect)
}
