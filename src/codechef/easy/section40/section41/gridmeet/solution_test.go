package main

import "testing"

func runSample(t *testing.T, K int, points [][]int, expect int64) {
	res := solve(K, points)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	K := 3
	points := [][]int{
		{1, 1},
		{2, 2},
		{2, 3},
		{3, 1},
		{4, 1},
	}
	var expect int64 = 3
	runSample(t, K, points, expect)
}

func TestSample2(t *testing.T) {
	K := 6
	points := [][]int{
		{-9, 8},
		{8, 5},
		{-5, -5},
		{2, 3},
		{11, 5},
		{1, 1},
		{-12, 11},
		{15, 0},
	}
	var expect int64 = 53
	runSample(t, K, points, expect)
}
