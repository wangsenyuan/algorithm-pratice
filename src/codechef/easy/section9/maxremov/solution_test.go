package main

import "testing"

func runSample(t *testing.T, N, K int, ops [][]int, expect int) {
	res := solve(N, K, ops)
	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", N, K, ops, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, K := 3, 2
	ops := [][]int{
		{2, 6},
		{4, 9},
		{1, 4},
	}
	expect := 3
	runSample(t, N, K, ops, expect)
}
