package main

import "testing"

func runSample(t *testing.T, values []int, E [][]int, expect int64) {
	res := solve(values, E)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	E := [][]int{
		{1, 2},
		{1, 3},
		{3, 4},
		{3, 5},
	}
	values := []int{4, 6, 8, 12, 15}
	var expect int64 = 29

	runSample(t, values, E, expect)
}
