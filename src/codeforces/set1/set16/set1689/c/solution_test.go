package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, expect int) {
	res := solve(n, E)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	E := [][]int{{1, 2}}
	expect := 0
	runSample(t, n, E, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	E := [][]int{
		{1, 2},
		{2, 3},
		{2, 4},
	}
	expect := 2
	runSample(t, n, E, expect)
}

func TestSample3(t *testing.T) {
	n := 7
	E := [][]int{
		{1, 2},
		{1, 5},
		{2, 3},
		{2, 4},
		{5, 6},
		{5, 7},
	}
	expect := 2
	runSample(t, n, E, expect)
}

func TestSample4(t *testing.T) {
	n := 15
	E := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
		{4, 6},
		{3, 7},
		{2, 8},
		{1, 9},
		{9, 10},
		{9, 11},
		{10, 12},
		{10, 13},
		{11, 14},
		{11, 15},
	}
	expect := 10
	runSample(t, n, E, expect)
}
