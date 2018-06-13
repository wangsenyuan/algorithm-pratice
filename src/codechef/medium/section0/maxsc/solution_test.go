package main

import "testing"

func runSample(t *testing.T, n int, matrix [][]int, expect int64) {
	res := solve(n, matrix)
	if res != expect {
		t.Errorf("sample %v, expect %d, but got %d", matrix, expect, res)
	}
}

func TestSample1(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	runSample(t, 3, matrix, 18)
}
