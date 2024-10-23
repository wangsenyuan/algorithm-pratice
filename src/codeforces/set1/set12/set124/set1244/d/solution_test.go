package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, c [][]int, expect int) {
	res, _ := solve(n, edges, c)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	c := [][]int{
		{3, 2, 3},
		{4, 3, 2},
		{3, 1, 3},
	}
	edges := [][]int{
		{1, 2},
		{2, 3},
	}
	expect := 6
	runSample(t, n, edges, c, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	c := [][]int{
		{3, 4, 2, 1, 2},
		{4, 2, 1, 5, 4},
		{5, 3, 2, 1, 1},
	}
	edges := [][]int{
		{1, 2},
		{3, 2},
		{4, 3},
		{5, 3},
	}
	expect := -1
	runSample(t, n, edges, c, expect)
}

func TestSample3(t *testing.T) {
	n := 5
	c := [][]int{
		{3, 4, 2, 1, 2},
		{4, 2, 1, 5, 4},
		{5, 3, 2, 1, 1},
	}
	edges := [][]int{
		{1, 2},
		{3, 2},
		{4, 3},
		{5, 4},
	}
	expect := 9
	runSample(t, n, edges, c, expect)
}
