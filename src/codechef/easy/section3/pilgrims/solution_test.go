package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, E []int64, expect int) {
	res := solve(n, edges, E)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	E := []int64{100, 20, 10, 50}
	edges := [][]int{
		{1, 2, 10},
		{2, 3, 20},
		{2, 4, 30},
		{4, 5, 40},
		{1, 6, 50},
	}
	expect := 2
	runSample(t, 6, edges, E, expect)
}
