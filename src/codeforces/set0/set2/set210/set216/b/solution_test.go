package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, expect int) {
	res := solve(n, edges)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	edges := [][]int{
		{1, 2},
		{2, 4},
		{5, 3},
		{1, 4},
	}
	expect := 1
	runSample(t, n, edges, expect)
}

func TestSample2(t *testing.T) {
	n := 6
	edges := [][]int{
		{1, 4},
		{3, 4},
	}
	expect := 0
	runSample(t, n, edges, expect)
}

func TestSample3(t *testing.T) {
	n := 6
	edges := [][]int{
		{1, 2},
		{2, 3},
		{3, 1},
		{4, 5},
		{5, 6},
		{6, 4},
	}
	expect := 2
	runSample(t, n, edges, expect)
}

func TestSample4(t *testing.T) {
	n := 4
	edges := [][]int{
		{3, 2},
		{4, 2},
		{4, 3},
	}
	expect := 2
	runSample(t, n, edges, expect)
}
