package main

import "testing"

func runSample(t *testing.T, G [][]int, k int, expect int64) {
	res := solve(G, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	G := [][]int{
		{2, 2, 3},
		{3, 4, 5},
		{4, 5, 5},
	}
	k := 4
	var expect int64 = 7
	runSample(t, G, k, expect)
}
