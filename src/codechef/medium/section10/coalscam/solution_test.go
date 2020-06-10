package main

import "testing"

func runSample(t *testing.T, N int, M1 int, M2 int, R [][]int, profit int64, total int64) {
	a, b := solve(N, M1, M2, R)

	if a != profit || b != total {
		t.Errorf("Sample expect %d %d, but got %d %d", profit, total, a, b)
	}
}

func TestSample1(t *testing.T) {
	N, M1, M2 := 3, 2, 1
	R := [][]int{
		{0, 1, 5},
		{1, 2, 4},
		{0, 1, 10},
	}
	var profit int64 = 10
	var total int64 = 14

	runSample(t, N, M1, M2, R, profit, total)
}

func TestSample2(t *testing.T) {
	N, M1, M2 := 3, 1, 1
	R := [][]int{
		{0, 1, 1},
		{0, 1, 3},
	}
	var profit int64 = -1
	var total int64 = -1

	runSample(t, N, M1, M2, R, profit, total)
}
