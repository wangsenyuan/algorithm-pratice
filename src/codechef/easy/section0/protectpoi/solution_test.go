package main

import "testing"

func runSample(t *testing.T, n int, k int, m int, snakes [][]int, expect int) {
	res := solve(n, k, m, snakes)
	if res != expect {
		t.Errorf("sampel: %d %d %d %v, should give answer %d, but got %d", n, k, m, snakes, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k, m := 7, 3, 7
	snakes := [][]int{
		{1, 1, 6, 1},
		{1, 2, 3, 2},
		{5, 2, 5, 2},
		{2, 4, 2, 6},
		{6, 2, 6, 4},
		{5, 6, 5, 7},
		{7, 1, 7, 4},
	}
	expect := 3
	runSample(t, n, k, m, snakes, expect)
}

func TestSample2(t *testing.T) {
	n, k, m := 7, 3, 7
	snakes := [][]int{
		{1, 1, 6, 1},
		{1, 2, 3, 2},
		{5, 2, 5, 2},
		{2, 6, 2, 6},
		{6, 2, 6, 4},
		{5, 6, 5, 7},
		{7, 1, 7, 4},
	}
	expect := -1
	runSample(t, n, k, m, snakes, expect)
}
