package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, X []int, expect int) {
	res := solve(n, E, X)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 8
	E := [][]int{
		{4, 7},
		{2, 5},
		{1, 6},
		{3, 6},
		{7, 2},
		{1, 7},
		{6, 8},
	}
	X := []int{5, 3}
	expect := -1
	runSample(t, n, E, X, expect)
}

func TestSample2(t *testing.T) {
	n := 8
	E := [][]int{
		{4, 7},
		{2, 5},
		{1, 6},
		{3, 6},
		{7, 2},
		{1, 7},
		{6, 8},
	}
	X := []int{6, 5, 7, 3}
	expect := 2
	runSample(t, n, E, X, expect)
}

func TestSample3(t *testing.T) {
	n := 3
	E := [][]int{
		{1, 2},
		{2, 3},
	}
	X := []int{2}
	expect := 1
	runSample(t, n, E, X, expect)
}

func TestSample4(t *testing.T) {
	n := 3
	E := [][]int{
		{1, 2},
		{1, 3},
	}
	X := []int{2, 3}
	expect := 2
	runSample(t, n, E, X, expect)
}
