package main

import "testing"

func runSample(t *testing.T, n, m int, A [][]int, expect int64) {
	res := solve(n, m, A)
	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", n, m, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 2, 3
	A := [][]int{
		{5, 15, 8},
		{3, 12, 10},
	}
	expect := int64(25)
	runSample(t, n, m, A, expect)
}
