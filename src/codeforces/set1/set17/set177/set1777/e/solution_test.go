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
		{1, 2, 3},
	}
	expect := 0
	runSample(t, n, edges, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	edges := [][]int{
		{1, 2, 10},
		{2, 3, 10},
		{3, 1, 10},
		{4, 5, 10},
	}
	expect := -1
	runSample(t, n, edges, expect)
}

func TestSample3(t *testing.T) {
	n := 4
	edges := [][]int{
		{1, 2, 10000},
		{2, 3, 20000},
		{1, 3, 30000},
		{4, 2, 500},
		{4, 3, 20},
	}
	expect := 20
	runSample(t, n, edges, expect)
}

func TestSample4(t *testing.T) {
	n := 4
	edges := [][]int{
		{1, 2, 10000},
		{2, 3, 20000},
		{1, 3, 30000},
		{4, 2, 5},
		{4, 3, 20},
	}
	expect := 5
	runSample(t, n, edges, expect)
}
