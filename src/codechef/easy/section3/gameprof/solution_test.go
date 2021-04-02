package main

import "testing"

func runSample(t *testing.T, k int, items [][]int, expect int64) {
	res := solve(k, items)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 2
	items := [][]int{
		{1, 3, 3},
		{1, 2, 5},
		{4, 7, 1},
	}
	var expect int64 = 4
	runSample(t, k, items, expect)
}
