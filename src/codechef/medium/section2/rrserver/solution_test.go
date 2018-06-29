package main

import "testing"

func runSample(t *testing.T, n int, m int, pairs [][]int, expect int) {
	res := solve(n, m, pairs)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", n, m, pairs, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 3, 2
	pairs := [][]int{
		{1, 2},
		{1, 3},
	}
	expect := 2
	runSample(t, n, m, pairs, expect)
}

func TestSample2(t *testing.T) {
	n, m := 5, 3
	pairs := [][]int{
		{1, 2},
		{2, 3},
		{4, 5},
	}
	runSample(t, n, m, pairs, 3)
}
