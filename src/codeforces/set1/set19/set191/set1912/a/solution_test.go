package main

import "testing"

func runSample(t *testing.T, x int, lists [][]int, expect int) {
	res := solve(x, lists)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	x := 1
	lists := [][]int{
		{-1, 2},
		{-2, 3},
		{-3, 4},
	}
	expect := 4
	runSample(t, x, lists, expect)
}

func TestSample2(t *testing.T) {
	x := 1
	lists := [][]int{
		{-1, -1, 4},
		{1, -3, -4, 8},
	}
	expect := 4
	runSample(t, x, lists, expect)
}
