package main

import "testing"

func TestSample(t *testing.T) {
	n := 4
	points := [][]int{
		{0, 0},
		{1000, 0},
		{1000, 2000},
		{0, 2000},
	}
	m := 2

	strips := [][]int{
		{1000, 10},
		{2000, 15},
	}

	res := solve(n, points, m, strips)

	if res != 50 {
		t.Errorf("Sample should give answer 50, but got %d", res)
	}
}
