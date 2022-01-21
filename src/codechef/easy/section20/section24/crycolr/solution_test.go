package main

import "testing"

func runSample(t *testing.T, n int, boxes [][]int, expect int) {
	res := solve(n, boxes)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	boxes := [][]int{
		{3, 0, 0},
		{0, 3, 0},
		{0, 0, 3},
	}
	expect := 0
	runSample(t, n, boxes, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	boxes := [][]int{
		{2, 1, 2},
		{1, 4, 0},
		{2, 0, 3},
	}
	expect := 3
	runSample(t, n, boxes, expect)
}
