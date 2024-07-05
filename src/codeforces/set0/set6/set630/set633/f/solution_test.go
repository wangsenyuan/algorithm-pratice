package main

import "testing"

func runSample(t *testing.T, a []int, edges [][]int, expect int) {
	res := solve(a, edges)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	edges := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{1, 5},
		{1, 6},
		{1, 7},
		{1, 8},
		{1, 9},
	}
	expect := 25
	runSample(t, a, edges, expect)
}

func TestSample2(t *testing.T) {
	a := []int{20, 10}
	edges := [][]int{
		{1, 2},
	}
	expect := 30
	runSample(t, a, edges, expect)
}
