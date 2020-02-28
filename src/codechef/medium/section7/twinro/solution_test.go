package main

import "testing"

func runSample(t *testing.T, N int, board [][]int, expect int64) {
	res := solve(N, board)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", N, board, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N := 4
	board := [][]int{
		{6, 0, 3, -1},
		{7, 4, 2, 4},
		{-3, 3, -2, 8},
		{13, 10, -1, -4},
	}
	var expect int64 = 56
	runSample(t, N, board, expect)
}
