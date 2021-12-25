package main

import "testing"

func runSample(t *testing.T, n, m int, G [][]int, expect int) {
	res := solve(n, m, G)
	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", n, m, G, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 3, 4
	G := [][]int{
		{1, 1, 1, 1},
		{-6, 1, 1, -4},
		{1, 1, 1, 1},
	}
	expect := 0
	runSample(t, n, m, G, expect)
}
