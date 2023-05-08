package main

import "testing"

func runSample(t *testing.T, n int, conds [][]int, expect int) {
	res := solve(n, conds)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	// 3, 2, 1, 4
	// 9, 12, 16, 32
	conds := [][]int{
		{3, 2, 3, 4},
		{1, 2, 4, 3},
		{1, 4, 2, 4},
	}
	expect := 69
	runSample(t, n, conds, expect)
}

func TestSample2(t *testing.T) {
	n := 8
	// 3, 2, 1, 4
	// 9, 12, 16, 32
	conds := [][]int{
		{5, 4, 2, 3},
		{6, 4, 5, 4},
		{1, 3, 5, 2},
		{6, 8, 2, 1},
		{3, 5, 3, 4},
		{3, 2, 2, 5},
		{6, 7, 4, 3},
	}
	expect := 359
	runSample(t, n, conds, expect)
}
