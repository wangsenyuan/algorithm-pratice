package main

import "testing"

func runSample(t *testing.T, k int, a []int, edges [][]int, expect int) {
	res := solve(k, a, edges)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 2
	a := []int{24, 12, 24, 6, 12}
	edges := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
		{2, 5},
	}
	expect := 288
	runSample(t, k, a, edges, expect)
}

func TestSample2(t *testing.T) {
	k := 3
	a := []int{24, 12, 24, 6, 12}
	edges := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
		{2, 5},
	}
	expect := 576
	runSample(t, k, a, edges, expect)
}
