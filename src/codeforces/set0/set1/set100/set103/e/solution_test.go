package main

import "testing"

func runSample(t *testing.T, n int, sets [][]int, X []int, expect int64) {
	res := solve(n, sets, X)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	sets := [][]int{
		{1},
		{2, 3},
		{3},
	}
	X := []int{10, 20, -3}
	var expect int64 = -3
	runSample(t, n, sets, X, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	sets := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
		{5, 1},
	}
	X := []int{1, -1, 1, -1, 1}
	var expect int64
	runSample(t, n, sets, X, expect)
}

func TestSample3(t *testing.T) {
	n := 5
	sets := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
		{5, 1},
	}
	X := []int{-1, 1, -1, 1, -1}
	var expect int64 = -1
	runSample(t, n, sets, X, expect)
}
