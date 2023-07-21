package main

import "testing"

func runSample(t *testing.T, H [][]int, V [][]int, expect int) {
	res := solve(H, V)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	H := [][]int{
		{2, 3, 1000000},
		{4, 0, 4},
		{3, 0, 1000000},
	}
	V := [][]int{
		{4, 0, 1},
		{2, 0, 5},
		{3, 1, 1000000},
	}
	expect := 7
	runSample(t, H, V, expect)
}
