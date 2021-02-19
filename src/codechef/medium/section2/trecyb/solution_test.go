package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, expect int64) {
	res := solve(n, E)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	E := [][]int{
		{1, 2},
		{1, 3},
	}
	var expect int64 = 2
	runSample(t, n, E, expect)
}
