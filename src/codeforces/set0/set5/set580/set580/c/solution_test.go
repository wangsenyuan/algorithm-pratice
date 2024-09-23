package main

import "testing"

func runSample(t *testing.T, m int, a []int, edges [][]int, expect int) {
	res := solve(m, a, edges)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	m := 1
	a := []int{1, 1, 0, 0}
	edges := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
	}
	expect := 2
	runSample(t, m, a, edges, expect)
}

func TestSample2(t *testing.T) {
	m := 1
	a := []int{1, 0, 1, 1, 0, 0, 0}
	edges := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
		{2, 5},
		{3, 6},
		{3, 7},
	}
	expect := 2
	runSample(t, m, a, edges, expect)
}
