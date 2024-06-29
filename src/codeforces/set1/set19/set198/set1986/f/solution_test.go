package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, expect int) {
	res := solve(n, edges)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	edges := [][]int{
		{1, 2},
	}
	expect := 0
	runSample(t, n, edges, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	edges := [][]int{
		{1, 2},
		{2, 3},
		{1, 3},
	}
	expect := 3
	runSample(t, n, edges, expect)
}

func TestSample3(t *testing.T) {
	n := 5
	edges := [][]int{
		{1, 2},
		{1, 3},
		{3, 4},
		{4, 5},
		{5, 3},
	}
	expect := 4
	runSample(t, n, edges, expect)
}

func TestSample4(t *testing.T) {
	n := 6
	edges := [][]int{
		{1, 2},
		{1, 3},
		{2, 3},
		{3, 4},
		{4, 5},
		{4, 6},
		{5, 6},
	}
	expect := 6
	runSample(t, n, edges, expect)
}

func TestSample5(t *testing.T) {
	n := 5
	edges := [][]int{
		{1, 2},
		{1, 3},
		{2, 3},
		{2, 4},
		{3, 5},
	}
	expect := 6
	runSample(t, n, edges, expect)
}

func TestSample6(t *testing.T) {
	n := 10
	edges := [][]int{
		{1, 2},
		{1, 3},
		{2, 3},
		{2, 4},
		{4, 5},
		{5, 6},
		{6, 7},
		{7, 4},
		{3, 8},
		{8, 9},
		{9, 10},
		{10, 8},
	}
	expect := 21
	runSample(t, n, edges, expect)
}
