package main

import "testing"

func runSample(t *testing.T, n int, m int, x int, y int, roads [][]int, taxis [][]int, expect int) {
	res := solve(n, m, x, y, roads, taxis)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 4, 4
	x, y := 1, 3
	roads := [][]int{
		{1, 2, 3},
		{1, 4, 1},
		{2, 4, 1},
		{2, 3, 5},
	}
	taxis := [][]int{
		{2, 7},
		{7, 2},
		{1, 2},
		{7, 7},
	}
	expect := 9
	runSample(t, n, m, x, y, roads, taxis, expect)
}
