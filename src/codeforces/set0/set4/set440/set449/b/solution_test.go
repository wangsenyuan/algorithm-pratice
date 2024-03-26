package main

import "testing"

func runSample(t *testing.T, n int, roads [][]int, s []int, y []int, expect int) {
	res := solve(n, roads, s, y)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	roads := [][]int{
		{1, 2, 1},
		{2, 3, 2},
		{1, 3, 3},
		{3, 4, 4},
		{1, 5, 5},
	}
	s := []int{3, 4, 5}
	y := []int{5, 5, 5}
	expect := 2
	runSample(t, n, roads, s, y, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	roads := [][]int{
		{1, 2, 2},
		{2, 1, 3},
	}
	s := []int{2, 2, 2}
	y := []int{1, 2, 3}
	expect := 2
	runSample(t, n, roads, s, y, expect)
}
