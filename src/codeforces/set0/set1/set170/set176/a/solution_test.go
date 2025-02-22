package main

import "testing"

func runSample(t *testing.T, k int, items [][][]int, expect int) {
	res := solve(k, items)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 10
	items := [][][]int{
		{
			{6, 5, 3},
			{7, 6, 5},
			{8, 6, 10},
		},
		{
			{10, 9, 0},
			{8, 6, 4},
			{10, 9, 3},
		},
		{
			{4, 3, 0},
			{8, 4, 12},
			{7, 2, 5},
		},
	}
	expect := 16
	runSample(t, k, items, expect)
}
