package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, expect int64) {
	res := solve(n, E)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	E := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
		{5, 1},
	}
	var expect int64 = 1
	runSample(t, n, E, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	E := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{1, 5},
		{2, 3},
		{2, 4},
		{2, 5},
		{3, 4},
		{3, 5},
		{4, 5},
	}
	
	var expect int64 = 12
	runSample(t, n, E, expect)
}
