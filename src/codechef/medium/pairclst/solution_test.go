package main

import (
	"testing"
)

func runSample(t *testing.T, N, M, K int, special []int, edges [][]int, expect int) {
	res := solve(N, M, K, special, edges)

	if res != expect {
		t.Errorf("Sample %d %d %d %v %v, expect %d, but got %d", N, M, K, special, edges, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, M, K := 5, 5, 3
	special := []int{1, 3, 5}
	edges := [][]int{
		{1, 2, 3},
		{2, 3, 4},
		{3, 4, 1},
		{4, 5, 8},
		{1, 5, 19},
	}
	expect := 7
	runSample(t, N, M, K, special, edges, expect)
}
