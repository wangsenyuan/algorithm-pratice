package main

import "testing"

func runSample(t *testing.T, d int, a []int, edges [][]int, expect int) {
	res := solve(d, a, edges)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	d := 1
	a := []int{2, 1, 3, 2}
	edges := [][]int{
		{1, 2},
		{1, 3},
		{3, 4},
	}
	expect := 8
	runSample(t, d, a, edges, expect)
}

func TestSample2(t *testing.T) {
	d := 0
	a := []int{1, 2, 3}
	edges := [][]int{
		{1, 2},
		{2, 3},
	}
	expect := 3
	runSample(t, d, a, edges, expect)
}

func TestSample3(t *testing.T) {
	d := 4
	a := []int{7, 8, 7, 5, 4, 6, 4, 10}
	edges := [][]int{
		{1, 6},
		{1, 2},
		{5, 8},
		{1, 3},
		{3, 5},
		{6, 7},
		{3, 4},
	}
	expect := 41
	runSample(t, d, a, edges, expect)
}
