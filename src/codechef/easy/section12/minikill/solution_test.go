package main

import "testing"

func runSample(t *testing.T, N int, edges [][]int, expect int) {
	res := solve(N, edges)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", N, edges, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N := 7
	edges := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
		{2, 5},
		{3, 6},
		{3, 7},
	}
	expect := 2
	runSample(t, N, edges, expect)
}

func TestSample2(t *testing.T) {
	N := 4
	edges := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
	}
	expect := 2
	runSample(t, N, edges, expect)
}

func TestSample3(t *testing.T) {
	N := 9
	edges := [][]int{
		{1, 2},
		{2, 4},
		{2, 3},
		{2, 5},
		{1, 6},
		{6, 7},
		{1, 8},
		{8, 9},
	}
	expect := 4
	runSample(t, N, edges, expect)
}

func TestSample4(t *testing.T) {
	N := 11
	edges := [][]int{
		{1, 2},
		{2, 4},
		{2, 3},
		{2, 5},
		{1, 6},
		{6, 7},
		{1, 8},
		{8, 9},
		{6, 10},
		{6, 11},
	}
	expect := 5
	runSample(t, N, edges, expect)
}
