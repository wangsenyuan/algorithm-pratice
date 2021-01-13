package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, expect int) {
	res := solve(n, E)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	E := [][]int{
		{1, 2},
		{1, 3},
	}
	expect := 2
	runSample(t, n, E, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	E := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
	}
	expect := 3
	runSample(t, n, E, expect)
}

func TestSample3(t *testing.T) {
	n := 8
	E := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{1, 5},
		{5, 6},
		{6, 7},
		{5, 8},
	}
	expect := 3
	runSample(t, n, E, expect)
}
