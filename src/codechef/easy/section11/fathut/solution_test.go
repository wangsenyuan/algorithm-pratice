package main

import "testing"

func runSample(t *testing.T, n int, F [][]int, expect int) {
	res := solve(n, F)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	F := [][]int{
		{5, 1, 1, 5},
		{4, 4, 2, 5},
		{2, 2, 2, 5},
		{3, 3, 2, 5},
	}
	expect := 1
	runSample(t, n, F, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	F := [][]int{
		{5, 1, 2, 5},
		{4, 4, 2, 5},
		{2, 2, 2, 5},
		{3, 3, 2, 5},
	}
	expect := 2
	runSample(t, n, F, expect)
}
