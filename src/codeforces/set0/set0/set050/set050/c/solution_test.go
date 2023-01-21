package main

import "testing"

func runSample(t *testing.T, ps [][]int, expect int) {
	res := solve(ps)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	ps := [][]int{
		{1, 1},
		{5, 1},
		{5, 3},
		{1, 3},
	}
	expect := 16
	runSample(t, ps, expect)
}

func TestSample2(t *testing.T) {
	ps := [][]int{
		{0, 2},
		{3, 0},
		{5, 3},
		{2, 5},
		{2, 2},
		{3, 3},
		{2, 2},
		{2, 2},
	}
	expect := 16
	runSample(t, ps, expect)
}
