package main

import "testing"

func runSample(t *testing.T, n, m int, B int, G [][]int, expect int) {
	res := solve(n, m, B, G)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m, B := 3, 3, 6
	G := [][]int{
		{1, 2, 3},
		{7, 5, 6},
		{2, 1, 2},
	}
	expect := 4
	runSample(t, n, m, B, G, expect)
}
