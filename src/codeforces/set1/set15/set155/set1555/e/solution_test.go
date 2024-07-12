package main

import "testing"

func runSample(t *testing.T, n int, m int, segs [][]int, expect int) {
	res := solve(n, m, segs)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 5, 12
	segs := [][]int{
		{1, 5, 5},
		{3, 4, 10},
		{4, 10, 6},
		{11, 12, 5},
		{10, 12, 3},
	}
	expect := 3
	runSample(t, n, m, segs, expect)
}

func TestSample2(t *testing.T) {
	n, m := 1, 10
	segs := [][]int{
		{1, 10, 23},
	}
	expect := 0
	runSample(t, n, m, segs, expect)
}
