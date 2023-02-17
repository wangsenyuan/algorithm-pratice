package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, expect int64) {
	res := solve(n, E)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	E := [][]int{
		{1, 2, 1},
		{1, 3, 2},
	}
	var expect int64 = 4
	runSample(t, n, E, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	E := [][]int{
		{1, 2, 2},
		{1, 3, 2},
	}
	var expect int64 = 2
	runSample(t, n, E, expect)
}

func TestSample3(t *testing.T) {
	n := 5
	E := [][]int{
		{1, 4, 4},
		{1, 2, 3},
		{3, 4, 4},
		{4, 5, 5},
	}
	var expect int64 = 14
	runSample(t, n, E, expect)
}

func TestSample4(t *testing.T) {
	n := 10
	E := [][]int{
		{10, 2, 3},
		{3, 8, 8},
		{4, 8, 9},
		{5, 8, 5},
		{3, 10, 7},
		{7, 8, 2},
		{5, 6, 6},
		{9, 3, 4},
		{1, 6, 3},
	}
	var expect int64 = 120
	runSample(t, n, E, expect)
}
