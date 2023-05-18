package main

import "testing"

func runSample(t *testing.T, n, k int, B []int, E [][]int, expect int) {
	res := solve(n, k, E, B)

	if res != int64(expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	B := []int{2, 8, 15, 9}
	n := 4
	k := 1
	E := [][]int{
		{1, 3},
		{2, 1},
		{4, 3},
		{3, 2},
	}
	expect := 25
	runSample(t, n, k, B, E, expect)
}
