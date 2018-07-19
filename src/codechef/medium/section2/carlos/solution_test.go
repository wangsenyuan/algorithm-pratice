package main

import "testing"

func runSample(t *testing.T, N, M, K int, T [][]int, W []int, expect int) {
	res := solve(N, M, K, T, W)

	if res != expect {
		t.Errorf("Sample %d %d %d %v %v, expect %d, but got %d", N, M, K, T, W, expect, res)
	}
}

func TestSample1(t *testing.T) {
	M, K, N := 3, 3, 3
	T := [][]int{
		{1, 2},
		{2, 3},
		{3, 1},
	}
	W := []int{3, 2, 1}
	expect := 2
	runSample(t, N, M, K, T, W, expect)
}

func TestSample2(t *testing.T) {
	M, K, N := 3, 1, 2
	T := [][]int{
		{2, 3},
	}
	W := []int{2, 1}
	expect := -1
	runSample(t, N, M, K, T, W, expect)
}

func TestSample3(t *testing.T) {
	M, K, N := 3, 1, 2
	T := [][]int{
		{2, 3},
	}
	W := []int{1, 3}
	expect := 0
	runSample(t, N, M, K, T, W, expect)
}
