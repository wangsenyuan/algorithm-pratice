package main

import "testing"

func runSample(t *testing.T, n int, P []int, E [][]int, expect int64) {
	res := solve(n, P, E)

	if res != expect {
		t.Errorf("Sample %d %v %v, expect %d, but got %d", n, P, E, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	P := []int{2, 1, 3}
	E := [][]int{
		{1, 2},
		{2, 3},
		{2, 4},
	}
	expect := int64(1)
	runSample(t, n, P, E, expect)
}
