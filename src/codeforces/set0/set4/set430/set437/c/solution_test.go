package main

import "testing"

func runSample(t *testing.T, vals []int, edges [][]int, expect int) {
	res := solve(vals, edges)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	vals := []int{10, 20, 30, 40}
	edges := [][]int{
		{1, 4},
		{1, 2},
		{2, 3},
	}
	expect := 40
	runSample(t, vals, edges, expect)
}

func TestSample2(t *testing.T) {
	vals := []int{100, 100, 100, 100}
	edges := [][]int{
		{1, 2},
		{2, 3},
		{2, 4},
		{3, 4},
	}
	expect := 400
	runSample(t, vals, edges, expect)
}

func TestSample3(t *testing.T) {
	vals := []int{40, 10, 20, 10, 20, 80, 40}
	edges := [][]int{
		{1, 5},
		{4, 7},
		{4, 5},
		{5, 2},
		{5, 7},
		{6, 4},
		{1, 6},
		{1, 3},
		{4, 3},
		{1, 4},
	}
	expect := 160
	runSample(t, vals, edges, expect)
}
