package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, expect [][]int) {
	_, res := solve(n, edges)

	a := getDiameter(n, edges, expect)
	b := getDiameter(n, edges, expect)

	if a != b {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	edges := [][]int{
		{1, 2},
		{2, 3},
	}
	expect := [][]int{{2, 4}}

	runSample(t, n, edges, expect)
}

func TestSample2(t *testing.T) {
	n := 2

	expect := [][]int{{1, 2}}

	runSample(t, n, nil, expect)
}

func TestSample3(t *testing.T) {
	n := 3
	edges := [][]int{
		{1, 2},
		{2, 3},
	}

	runSample(t, n, edges, nil)
}

func TestSample4(t *testing.T) {
	n := 5
	extra := [][]int{
		{3, 5},
		{2, 5},
		{1, 5},
		{4, 5},
	}

	runSample(t, n, nil, extra)
}
