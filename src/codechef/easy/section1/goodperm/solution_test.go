package main

import "testing"

func runSample(t *testing.T, N, K int, A []int, expect int) {
	res := solve(N, K, A)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", N, K, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, K := 3, 1
	A := []int{2, 0, 0}
	expect := 2
	runSample(t, N, K, A, expect)
}
