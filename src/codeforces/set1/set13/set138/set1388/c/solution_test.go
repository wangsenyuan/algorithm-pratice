package main

import "testing"

func runSample(t *testing.T, a []int, h []int, edges [][]int, expect bool) {
	res := solve(a, h, edges)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 0, 1, 1, 0, 1, 0}
	h := []int{4, 0, 0, -1, 0, -1, 0}
	edges := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{3, 5},
		{3, 6},
		{3, 7},
	}
	expect := true
	runSample(t, a, h, edges, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 2, 5, 2, 1}
	h := []int{-11, -2, -6, -2, -1}
	edges := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{3, 5},
	}
	expect := true
	runSample(t, a, h, edges, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 1, 1, 1}
	h := []int{4, 1, -3, -1}
	edges := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
	}
	expect := false
	runSample(t, a, h, edges, expect)
}

func TestSample4(t *testing.T) {
	a := []int{3, 3, 7}
	h := []int{13, 1, 4}
	edges := [][]int{
		{1, 2},
		{1, 3},
	}
	expect := false
	runSample(t, a, h, edges, expect)
}
