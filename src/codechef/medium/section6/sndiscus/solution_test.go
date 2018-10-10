package main

import "testing"

func runSample(t *testing.T, n int, snakes [][]int, expect int) {
	res := solve(n, snakes)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, snakes, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	snakes := [][]int{
		{1, 1, 1, 3},
		{1, 1, 3, 1},
	}
	expect := 0
	runSample(t, n, snakes, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	snakes := [][]int{
		{1, 1, 1, 3},
		{2, 1, 3, 1},
	}
	expect := 1
	runSample(t, n, snakes, expect)
}

func TestSample3(t *testing.T) {
	n := 3
	snakes := [][]int{
		{1, 1, 1, 3},
		{2, 1, 3, 1},
		{2, 2, 2, 2},
	}
	expect := 1
	runSample(t, n, snakes, expect)
}
