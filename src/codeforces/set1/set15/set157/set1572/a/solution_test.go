package main

import "testing"

func runSample(t *testing.T, n int, reqs [][]int, expect int) {
	res := solve(n, reqs)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	reqs := [][]int{
		{2},
		{},
		{1, 4},
		{2},
	}
	expect := 2
	runSample(t, n, reqs, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	reqs := [][]int{
		{5},
		{1},
		{2},
		{3},
		{4},
	}
	expect := -1
	runSample(t, n, reqs, expect)
}
