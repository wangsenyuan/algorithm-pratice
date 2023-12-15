package main

import "testing"

func runSample(t *testing.T, n int, k int, c int, E [][]int, expect int) {
	res := solve(n, k, c, E)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k, c := 3, 2, 3
	E := [][]int{
		{2, 1},
		{3, 1},
	}
	expect := 2
	runSample(t, n, k, c, E, expect)
}

func TestSample2(t *testing.T) {
	n, k, c := 5, 4, 1
	E := [][]int{
		{2, 1},
		{4, 2},
		{5, 4},
		{3, 4},
	}
	expect := 12
	runSample(t, n, k, c, E, expect)
}

func TestSample3(t *testing.T) {
	n, k, c := 6, 5, 3
	E := [][]int{
		{4, 1},
		{6, 1},
		{2, 6},
		{5, 1},
		{3, 2},
	}
	expect := 17
	runSample(t, n, k, c, E, expect)
}

func TestSample4(t *testing.T) {
	n, k, c := 10, 6, 4
	E := [][]int{
		{1, 3},
		{1, 9},
		{9, 7},
		{7, 6},
		{6, 4},
		{9, 2},
		{2, 8},
		{8, 5},
		{5, 10},
	}
	expect := 32
	runSample(t, n, k, c, E, expect)
}
