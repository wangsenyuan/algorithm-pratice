package main

import "testing"

func runSample(t *testing.T, p, q, c int, M [][]int, expect int) {
	res := solve(p, q, c, M)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	p, q, c := 5, 2, 2
	M := [][]int{
		{3, 4},
		{1, 5},
		{4, 4},
		{5, 0},
		{2, 2},
		{5, 0},
	}
	expect := 4
	runSample(t, p, q, c, M, expect)
}
