package main

import "testing"

func runSample(t *testing.T, R, C, D int, G [][]int, expect int) {
	res := solve(R, C, D, G)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	D, R, C := 2, 0, 4
	G := [][]int{
		{0, 1, 5, 1, 4},
	}
	expect := 6
	runSample(t, R, C, D, G, expect)
}

func TestSample2(t *testing.T) {
	D, R, C := 2, 4, 3
	G := [][]int{
		{0, 7, 8, 5, 9, 1},
		{1, 6, 8, 4, 6, 2},
		{5, 4, 2, 5, 0, 3},
		{5, 2, 0, 6, 8, 8},
		{3, 5, 3, 3, 8, 4},
	}
	expect := 4
	runSample(t, R, C, D, G, expect)
}
