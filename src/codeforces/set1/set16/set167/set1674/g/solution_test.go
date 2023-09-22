package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, expect int) {
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
	expect := 2
	runSample(t, n, E, expect)
}

func TestSample2(t *testing.T) {
	n := 7
	E := [][]int{
		{7, 1},
		{1, 3},
		{6, 2},
		{2, 3},
		{7, 2},
		{2, 4},
		{7, 3},
		{6, 3},
	}
	expect := 3
	runSample(t, n, E, expect)
}

func TestSample3(t *testing.T) {
	n := 8
	E := [][]int{
		{1, 3},
		{2, 3},
		{4, 5},
		{3, 5},
		{3, 6},
		{5, 6},
		{5, 7},
		{8, 7},
	}
	expect := 3
	runSample(t, n, E, expect)
}
