package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, b int, e int, expect int) {
	res := solve(n, edges, b, e)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	edges := [][]int{
		{1, 2, 1},
		{2, 3, 1},
		{5, 2, 2},
		{2, 4, 2},
		{4, 6, 2},
		{3, 6, 3},
	}
	b, e := 1, 3
	expect := 1
	runSample(t, n, edges, b, e, expect)
}

func TestSample2(t *testing.T) {
	n := 6
	edges := [][]int{
		{1, 2, 1},
		{2, 3, 1},
		{5, 2, 2},
		{2, 4, 2},
		{4, 6, 2},
		{3, 6, 3},
	}
	b, e := 1, 6
	expect := 2
	runSample(t, n, edges, b, e, expect)
}

func TestSample3(t *testing.T) {
	n := 4
	edges := [][]int{
		{1, 2, 1},
		{1, 3, 1},
		{4, 1, 1},
	}
	b, e := 2, 3
	expect := 1
	runSample(t, n, edges, b, e, expect)
}

func TestSample4(t *testing.T) {
	n := 6
	edges := [][]int{
		{1, 2, 43},
		{1, 3, 34},
		{4, 6, 43},
		{6, 3, 43},
		{2, 3, 43},
		{5, 3, 43},
		{4, 5, 43},
	}
	b, e := 1, 6
	expect := 1
	runSample(t, n, edges, b, e, expect)
}
