package main

import "testing"

func runSample(t *testing.T, a []int, edges [][]int, expect int) {
	res := solve(a, edges)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, -2, 2, 1}
	edges := [][]int{
		{1, 2},
		{3, 2},
		{2, 4},
	}
	expect := 3
	runSample(t, a, edges, expect)
}

func TestSample2(t *testing.T) {
	a := []int{-2, -5}
	edges := [][]int{
		{1, 2},
	}
	expect := 0
	runSample(t, a, edges, expect)
}

func TestSample3(t *testing.T) {
	a := []int{-2, 4, -2, 3, 3, 2, -1}
	edges := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{3, 5},
		{4, 6},
		{4, 7},
	}
	expect := 9
	runSample(t, a, edges, expect)
}
