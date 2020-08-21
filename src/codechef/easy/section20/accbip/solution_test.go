package main

import "testing"

func runSample(t *testing.T, N, C, K int, L [][]int, V []int, expect int64) {
	res := solve(N, C, K, L, V)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, C, K := 7, 2, 13
	L := [][]int{
		{1, 10, 1},
		{1, 14, 2},
		{6, 4, 1},
		{2, 2, 1},
		{0, 12, 2},
		{2, 11, 2},
		{0, 6, 1},
	}
	V := []int{8, 10}
	var expect int64 = 2
	runSample(t, N, C, K, L, V, expect)
}

func TestSample2(t *testing.T) {
	N, C, K := 6, 1, 20
	L := [][]int{
		{1, 5, 1},
		{2, 11, 1},
		{4, 0, 1},
		{6, 8, 1},
		{0, 11, 1},
		{3, 3, 1},
	}
	V := []int{9}
	var expect int64 = 4
	runSample(t, N, C, K, L, V, expect)
}
