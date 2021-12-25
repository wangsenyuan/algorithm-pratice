package main

import "testing"

func runSample(t *testing.T, N, K int, E [][]int, A []int, expect int64) {
	res := solve(N, len(E), len(A), K, E, A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, K := 2, 2
	E := [][]int{
		{0, 1},
		{1, 2},
		{2, 0},
	}
	A := []int{1, 2}
	var expect int64 = 4
	runSample(t, N, K, E, A, expect)
}

func TestSample2(t *testing.T) {
	N, K := 2, 2
	E := [][]int{
		{0, 1},
		{1, 2},
	}
	A := []int{1, 2}
	var expect int64 = 6
	runSample(t, N, K, E, A, expect)
}

func TestSample3(t *testing.T) {
	N, K := 6, 3
	E := [][]int{
		{0, 1},
		{0, 2},
		{0, 4},
		{1, 3},
		{1, 4},
		{2, 5},
		{2, 6},
	}
	A := []int{1, 2, 3, 5, 6}
	var expect int64 = 8
	runSample(t, N, K, E, A, expect)
}
