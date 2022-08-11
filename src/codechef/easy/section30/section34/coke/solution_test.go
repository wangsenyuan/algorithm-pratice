package main

import "testing"

func runSample(t *testing.T, cakes [][]int, M int, K int, L int, R int, expect int) {
	res := solve(cakes, M, K, L, R)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	M, K, L, R := 2, 5, 4, 6
	cakes := [][]int{
		{1, 6},
		{2, 8},
		{8, 10},
	}
	expect := 8
	runSample(t, cakes, M, K, L, R, expect)
}
