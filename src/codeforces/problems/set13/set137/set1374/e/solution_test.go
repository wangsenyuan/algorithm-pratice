package main

import "testing"

func runSample(t *testing.T, n int, k int, books [][]int, expect int64) {
	res := solve(n, k, books)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", n, k, books, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 8, 4
	books := [][]int{
		{7, 1, 1},
		{2, 1, 1},
		{4, 0, 1},
		{8, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
		{1, 0, 1},
		{3, 0, 0},
	}
	var expect int64 = 18
	runSample(t, n, k, books, expect)
}

func TestSample2(t *testing.T) {
	n, k := 5, 2
	books := [][]int{
		{6, 0, 0},
		{9, 0, 0},
		{1, 0, 1},
		{2, 1, 1},
		{5, 1, 0},
	}
	var expect int64 = 8
	runSample(t, n, k, books, expect)
}
func TestSample3(t *testing.T) {
	n, k := 5, 3
	books := [][]int{
		{3, 0, 0},
		{2, 1, 0},
		{3, 1, 0},
		{5, 0, 1},
		{3, 0, 1},
	}
	var expect int64 = -1
	runSample(t, n, k, books, expect)
}
