package main

import "testing"

func runSample(t *testing.T, n int, X int64, G [][]int64, expect int64) {
	res := solve(n, X, G)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	var X int64 = 10
	G := [][]int64{
		{2, 2, 3},
		{2, 11, 1},
		{2, 3, 1},
	}
	var expect int64 = 2
	runSample(t, n, X, G, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	var X int64 = 50
	G := [][]int64{
		{2, 3, 4, 7, 12},
		{4, 5, 6, 3, 5},
		{5, 5, 6, 9, 1},
		{1, 2, 9, 6, 7},
		{5, 6, 7, 4, 7},
	}
	var expect int64 = 62
	runSample(t, n, X, G, expect)
}
