package main

import "testing"

func runSample(t *testing.T, N int, K int, A []int, expect int) {
	res := solve(N, K, A)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", N, K, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 5, 2, []int{3, 6, 4, 7, 2}, 13)
}

func TestSample2(t *testing.T) {
	runSample(t, 5, 3, []int{3, -5, 6, 3, 10}, 10)
}
