package main

import "testing"

func runSample(t *testing.T, colors []int, edges [][]int, expect int) {
	res := solve(colors, edges)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	colors := []int{1, 2, 1}
	edges := [][]int{
		{1, 2},
		{2, 3},
	}
	expect := 1
	runSample(t, colors, edges, expect)
}

func TestSample2(t *testing.T) {
	colors := []int{2, 1, 2, 1, 2}
	edges := [][]int{
		{1, 2},
		{1, 3},
		{3, 4},
		{4, 5},
	}
	expect := 3
	runSample(t, colors, edges, expect)
}

func TestSample3(t *testing.T) {
	colors := []int{2, 2, 2, 2}
	edges := [][]int{
		{3, 1},
		{3, 2},
		{3, 4},
	}
	expect := 3
	runSample(t, colors, edges, expect)
}
