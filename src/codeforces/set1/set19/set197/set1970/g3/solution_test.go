package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, c int, expect int) {
	res := solve(n, edges, c)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, c := 4, 5
	edges := [][]int{
		{4, 3},
		{2, 3},
		{2, 4},
		{1, 2},
		{4, 1},
		{3, 1},
	}
	expect := -1
	runSample(t, n, edges, c, expect)
}

func TestSample2(t *testing.T) {
	n, c := 6, 2
	edges := [][]int{
		{1, 4},
		{2, 5},
		{3, 6},
		{1, 5},
		{3, 5},
		{6, 5},
	}
	expect := 20
	runSample(t, n, edges, c, expect)
}

func TestSample3(t *testing.T) {
	n, c := 6, 7
	edges := [][]int{
		{1, 4},
		{2, 5},
		{3, 6},
		{3, 5},
		{6, 5},
	}
	expect := 25
	runSample(t, n, edges, c, expect)
}

func TestSample4(t *testing.T) {
	n, c := 7, 4
	edges := [][]int{
		{1, 4},
		{3, 6},
		{3, 5},
		{6, 5},
		{2, 7},
	}
	expect := 33
	runSample(t, n, edges, c, expect)
}
