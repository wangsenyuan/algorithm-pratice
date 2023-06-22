package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, k int, expect int64) {
	res := solve(n, E, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 4, 7
	E := [][]int{
		{4, 1, 3},
		{1, 2, 5},
		{2, 3, 8},
		{2, 4, 1},
		{3, 4, 4},
	}
	var expect int64 = 1
	runSample(t, n, E, k, expect)
}

func TestSample2(t *testing.T) {
	n, k := 4, 5
	E := [][]int{
		{1, 2, 1},
		{1, 3, 1},
		{1, 4, 2},
		{2, 4, 1},
		{4, 3, 1},
		{3, 2, 1},
	}
	var expect int64 = 3
	runSample(t, n, E, k, expect)
}

func TestSample3(t *testing.T) {
	n, k := 3, 10
	E := [][]int{
		{1, 2, 8},
		{1, 3, 10},
	}
	var expect int64
	runSample(t, n, E, k, expect)
}

func TestSample4(t *testing.T) {
	n, k := 5, 15
	E := [][]int{
		{1, 2, 17},
		{3, 1, 15},
		{2, 3, 10},
		{1, 4, 14},
		{2, 5, 8},
	}
	var expect int64
	runSample(t, n, E, k, expect)
}
