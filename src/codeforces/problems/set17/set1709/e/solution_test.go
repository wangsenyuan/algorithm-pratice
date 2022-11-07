package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, A []int, expect int) {
	res := solve(n, E, A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	A := []int{3, 2, 1, 3, 2, 1}
	E := [][]int{
		{4, 5},
		{3, 4},
		{1, 4},
		{2, 1},
		{6, 1},
	}
	expect := 2
	runSample(t, n, E, A, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	A := []int{2, 1, 1, 1}
	E := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
	}
	expect := 0
	runSample(t, n, E, A, expect)
}

func TestSample3(t *testing.T) {
	n := 5
	A := []int{2, 2, 2, 2, 2}
	E := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
	}
	expect := 2
	runSample(t, n, E, A, expect)
}

func TestSample4(t *testing.T) {
	n := 10
	A := []int{7, 4, 6, 7, 6, 6, 7, 5, 7, 5}
	E := [][]int{
		{8, 7},
		{4, 5},
		{9, 6},
		{2, 5},
		{4, 8},
		{9, 10},
		{4, 3},
		{9, 4},
		{1, 8},
	}
	expect := 1
	runSample(t, n, E, A, expect)
}
