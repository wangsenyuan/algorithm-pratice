package main

import "testing"

func runSample(t *testing.T, m int64, intervals [][]int64, expect int64) {
	res := solve(m, intervals)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	var m int64 = 9
	intervals := [][]int64{
		{0, 2},
		{3, 4},
		{5, 7},
	}
	var expect int64 = 3
	runSample(t, m, intervals, expect)
}
