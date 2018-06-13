package main

import "testing"

func TestSampleOne(t *testing.T) {
	n, m := 5, 2
	ss := [][]int{
		{0, 1, 2},
		{2, 3, 4},
	}
	res := solve(n, m, ss)
	if res != 3 {
		t.Errorf("Sample one should give result 3, but got %d", res)
	}
}

func TestSampleTwo(t *testing.T) {
	n, m := 4, 3
	ss := [][]int{
		{0, 1},
		{1, 2},
		{2, 3},
	}
	res := solve(n, m, ss)
	if res != 4 {
		t.Errorf("Sample two should give result 4, but got %d", res)
	}
}
