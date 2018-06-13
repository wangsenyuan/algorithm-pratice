package main

import "testing"

func runSample(t *testing.T, m, n int, mat [][]int, expect int) {
	res := solve(m, n, mat)
	if res != expect {
		t.Errorf("sample: %d %d %v should give answer %d, but got %d", m, n, mat, expect, res)
	}
}

func TestSample1(t *testing.T) {
	m, n := 3, 3
	mat := [][]int{
		{1, 1, 1},
		{1, 3, -1},
		{1, 1, 1},
	}
	expect := 6
	runSample(t, m, n, mat, expect)
}

func TestSample2(t *testing.T) {
	m, n := 4, 5
	mat := [][]int{
		{1, -2, 3, -4, 5},
		{-5, 4, -3, 2, -1},
		{1, -3, 5, -7, 9},
		{-9, 7, -5, 3, -1},
	}
	expect := 19
	runSample(t, m, n, mat, expect)
}
