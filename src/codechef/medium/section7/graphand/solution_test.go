package main

import "testing"

func runSample(tc *testing.T, n int, m int, E [][]int, s, t, k int, expect int64) {
	res := solve(n, m, E, s, t, k)

	if res != expect {
		tc.Errorf("Sample %d %d %v %d %d %d, expect %d, but got %d", n, m, E, s, t, k, expect, res)
	}
}

func TestSample1(tc *testing.T) {
	n, m := 3, 3
	E := [][]int{
		{1, 2, 2, 3},
		{2, 3, 1, 7},
		{1, 3, 2, 2},
	}
	s, t, k := 1, 3, 3
	expect := int64(3)
	runSample(tc, n, m, E, s, t, k, expect)
}
