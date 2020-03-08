package main

import "testing"

func runSample(t *testing.T, N, M int, E [][]int, expect int) {
	res := solve(N, M, E)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", N, M, E, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, M := 5, 5
	E := [][]int{
		{5, 1},
		{5, 2},
		{1, 2},
		{2, 3},
		{2, 4},
	}
	expect := 1
	runSample(t, N, M, E, expect)
}

func TestSample2(t *testing.T) {
	N, M := 5, 6
	E := [][]int{
		{4, 5},
		{4, 1},
		{4, 2},
		{4, 3},
		{5, 1},
		{5, 2},
	}
	expect := 4
	runSample(t, N, M, E, expect)
}

func TestSample3(t *testing.T) {
	N, M := 5, 5
	E := [][]int{
		{3, 4},
		{3, 5},
		{3, 1},
		{3, 2},
		{4, 2},
	}
	expect := 2
	runSample(t, N, M, E, expect)
}

func TestSample4(t *testing.T) {
	N, M := 4, 1
	E := [][]int{
		{3, 4},
	}
	expect := -1
	runSample(t, N, M, E, expect)
}

func TestSample5(t *testing.T) {
	N, M := 6, 6
	E := [][]int{
		{1, 2},
		{2, 3},
		{3, 1},
		{4, 5},
		{5, 6},
		{6, 4},
	}
	expect := -1
	runSample(t, N, M, E, expect)
}
