package main

import "testing"

func runSample(t *testing.T, p, n int, E [][]int, W []int, expect int) {
	res := solve(p, n, E, W)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, p := 4, 2
	W := []int{7, 4, 3, 1}
	E := [][]int{
		{1, 2, 21},
		{3, 2, 6},
		{1, 3, 8},
		{2, 4, 11},
	}
	expect := 4
	runSample(t, p, n, E, W, expect)
}

func TestSample2(t *testing.T) {
	n, p := 4, 10
	W := []int{1, 2, 10, 1}
	E := [][]int{
		{1, 2, 20},
		{2, 4, 30},
		{1, 3, 25},
		{3, 4, 89},
	}
	expect := 24
	runSample(t, p, n, E, W, expect)
}
