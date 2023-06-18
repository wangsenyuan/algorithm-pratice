package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, expect int64) {
	res := solve(n, E)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	E := [][]int{
		{1, 2},
		{2, 3},
		{1, 3},
	}
	var expect int64 = 6
	runSample(t, n, E, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	E := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 2},
	}
	var expect int64 = 11
	runSample(t, n, E, expect)
}

func TestSample3(t *testing.T) {
	n := 5
	E := [][]int{
		{1, 2},
		{2, 3},
		{1, 3},
		{2, 5},
		{4, 3},
	}
	var expect int64 = 18
	runSample(t, n, E, expect)
}

func TestSample4(t *testing.T) {
	n := 11
	E := [][]int{
		{9, 1},
		{1, 2},
		{7, 1},
		{4, 10},
		{8, 5},
		{4, 6},
		{4, 2},
		{7, 5},
		{2, 5},
		{3, 1},
		{11, 6},
	}
	var expect int64 = 96
	runSample(t, n, E, expect)
}
