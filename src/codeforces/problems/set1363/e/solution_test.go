package main

import "testing"

func runSample(t *testing.T, n int, X [][]int, E [][]int, expect int64) {
	res := solve(n, X, E)

	if res != expect {
		t.Errorf("Sample %d %v %v, expect %d, but got %d", n, X, E, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	X := [][]int{
		{1, 0, 1},
		{20, 1, 0},
		{300, 0, 1},
		{4000, 0, 0},
		{50000, 1, 0},
	}
	E := [][]int{
		{1, 2},
		{2, 3},
		{2, 4},
		{1, 5},
	}
	var expect int64 = 4

	runSample(t, n, X, E, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	X := [][]int{
		{10000, 0, 1},
		{2000, 1, 0},
		{300, 0, 1},
		{40, 0, 0},
		{1, 1, 0},
	}
	E := [][]int{
		{1, 2},
		{2, 3},
		{2, 4},
		{1, 5},
	}
	var expect int64 = 24000

	runSample(t, n, X, E, expect)
}

func TestSample3(t *testing.T) {
	n := 2
	X := [][]int{
		{109, 0, 1},
		{205, 0, 1},
	}
	E := [][]int{
		{1, 2},
	}
	var expect int64 = -1

	runSample(t, n, X, E, expect)
}

func TestSample4(t *testing.T) {
	n := 8
	X := [][]int{
		{82, 0, 0},
		{3, 1, 1},
		{53, 0, 0},
		{5, 0, 0},
		{81, 0, 1},
		{56, 1, 0},
		{99, 1, 0},
		{87, 0, 1},
	}
	E := [][]int{
		{5, 7},
		{2, 8},
		{4 ,7},
		{4, 1},
		{3, 2},
		{2 ,5},
		{8 ,6},
	}
	var expect int64 = 16

	runSample(t, n, X, E, expect)
}
