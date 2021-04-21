package main

import "testing"

func runSample(t *testing.T, K int, N int, L int, R int, arrs [][]int, expect int) {
	res := solve(K, N, L, R, arrs)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	K, N, L, R := 4, 3, 1, 9
	arrs := [][]int{
		{1, 1, 2},
		{1, 2, 3},
		{1, 2, 4},
		{1, 2, 2},
	}
	expect := 1
	runSample(t, K, N, L, R, arrs, expect)
}
