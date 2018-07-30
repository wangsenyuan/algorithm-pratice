package main

import "testing"

func runSample(t *testing.T, n, m int, G [][]int, expect int) {
	res := solve(n, m, G)
	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", n, m, G, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 3, 3
	G := [][]int{
		{2, 2, 1},
		{3, 4, 3},
		{2, 3, 2},
	}
	expect := 1
	runSample(t, n, m, G, expect)
}

func TestSample2(t *testing.T) {
	n, m := 3, 4
	G := [][]int{
		{1, 2, 1, 1},
		{2, 3, 3, 2},
		{2, 2, 2, 1},
	}
	expect := 1
	runSample(t, n, m, G, expect)
}
