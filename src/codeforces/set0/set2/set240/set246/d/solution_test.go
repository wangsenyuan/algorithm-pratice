package main

import "testing"

func runSample(t *testing.T, c []int, edges [][]int, expect int) {
	res := solve(c, edges)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	c := []int{1, 1, 2, 3, 5, 8}
	edges := [][]int{
		{1, 2},
		{3, 2},
		{1, 4},
		{4, 3},
		{4, 5},
		{4, 6},
	}
	expect := 3
	runSample(t, c, edges, expect)
}

func TestSample2(t *testing.T) {
	c := []int{4, 2, 5, 2, 4}
	edges := [][]int{
		{1, 2},
		{2, 3},
		{3, 1},
		{5, 3},
		{5, 4},
		{3, 4},
	}
	expect := 2
	runSample(t, c, edges, expect)
}
