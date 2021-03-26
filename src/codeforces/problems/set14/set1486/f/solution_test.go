package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, paths [][]int, expect int64) {
	res := solve(n, E, paths)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	E := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{3, 5},
	}

	paths := [][]int{
		{2, 3},
		{2, 4},
		{3, 4},
		{3, 5},
	}
	var expect int64 = 2
	runSample(t, n, E, paths, expect)
}

func TestSample2(t *testing.T) {
	n := 1
	E := [][]int{

	}

	paths := [][]int{
		{1, 1},
		{1, 1},
		{1, 1},
	}
	var expect int64 = 3
	runSample(t, n, E, paths, expect)
}

func TestSample3(t *testing.T) {
	n := 5
	E := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{3, 5},
	}

	paths := [][]int{
		{2, 3},
		{2, 4},
		{3, 4},
		{3, 5},
		{1, 1},
		{1, 2},
	}
	var expect int64 = 7
	runSample(t, n, E, paths, expect)
}

func TestSample4(t *testing.T) {
	n := 10
	E := [][]int{
		{10, 1},
		{8, 5},
		{5, 3},
		{10, 8},
		{10, 7},
		{3, 4},
		{9, 5},
		{2, 8},
		{6, 3},
	}

	paths := [][]int{
		{2, 1},
		{1, 6},
		{2, 3},
		{3, 9},
		{10, 10},
		{7, 10},
		{9, 4},
		{4, 5},
		{1, 2},
		{7, 3},
	}
	var expect int64 = 8
	runSample(t, n, E, paths, expect)
}
