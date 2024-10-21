package main

import "testing"

func runSample(t *testing.T, a []int, n int, traps [][]int, x int, expect int) {
	res := solve(a, n, traps, x)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, x := 6, 14
	a := []int{1, 2, 3, 4, 5}
	traps := [][]int{
		{1, 5, 2},
		{1, 2, 5},
		{2, 3, 5},
		{3, 5, 3},
	}
	expect := 3
	runSample(t, a, n, traps, x, expect)
}
