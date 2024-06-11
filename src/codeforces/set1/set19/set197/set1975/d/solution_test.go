package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, P []int, expect int) {
	res := solve(n, edges, P)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	edges := [][]int{
		{1, 2},
	}
	P := []int{1, 2}
	expect := 2
	runSample(t, n, edges, P, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	edges := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{1, 5},
	}
	P := []int{1, 2}
	expect := 8
	runSample(t, n, edges, P, expect)
}

func TestSample3(t *testing.T) {
	n := 8
	edges := [][]int{
		{7, 1},
		{1, 5},
		{1, 8},
		{8, 3},
		{7, 2},
		{8, 6},
		{3, 4},
	}
	P := []int{5, 4}
	expect := 13
	runSample(t, n, edges, P, expect)
}

func TestSample4(t *testing.T) {
	n := 5
	edges := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
	}
	P := []int{1, 5}
	expect := 8
	runSample(t, n, edges, P, expect)
}
