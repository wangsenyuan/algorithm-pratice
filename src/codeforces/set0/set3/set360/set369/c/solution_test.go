package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, expect int) {
	res := solve(n, E)

	if len(res) != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	E := [][]int{{1, 2, 2}, {2, 3, 2}, {3, 4, 2}, {4, 5, 2}}
	expect := 1
	runSample(t, n, E, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	E := [][]int{
		{1, 2, 1},
		{2, 3, 2},
		{2, 4, 1},
		{4, 5, 1}}
	expect := 1
	runSample(t, n, E, expect)
}

func TestSample3(t *testing.T) {
	n := 10
	E := [][]int{
		{7, 5, 1},
		{2, 1, 2},
		{8, 7, 2},
		{2, 4, 1},
		{4, 5, 2},
		{9, 5, 1},
		{3, 2, 2},
		{2, 10, 1},
		{6, 5, 2},
	}
	expect := 3
	runSample(t, n, E, expect)
}

func TestSample4(t *testing.T) {
	n := 5
	E := [][]int{
		{1, 5, 1},
		{5, 4, 2},
		{4, 3, 1},
		{3, 2, 2},
	}
	expect := 1
	runSample(t, n, E, expect)
}
