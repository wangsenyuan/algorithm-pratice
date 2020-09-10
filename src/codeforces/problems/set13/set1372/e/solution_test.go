package main

import "testing"

func runSample(t *testing.T, n, m int, segs [][][]int, expect int) {
	res := solve(n, m, segs)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 4, 5
	segs := [][][]int{
		{
			{1, 2},
			{2, 5},
		},
		{
			{1, 3},
			{4, 5},
		},
		{
			{1, 1},
			{2, 4},
			{5, 5},
		},
		{
			{1, 1},
			{2, 2},
			{3, 5},
		},
	}
	expect := 36
	runSample(t, n, m, segs, expect)
}
