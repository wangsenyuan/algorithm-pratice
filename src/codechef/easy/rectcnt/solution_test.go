package main

import "testing"

func TestSample(t *testing.T) {
	n := 6
	coords := [][]int{
		[]int{7, 1},
		[]int{3, 5},
		[]int{3, 1},
		[]int{1, 5},
		[]int{1, 1},
		[]int{7, 5},
	}
	res := solve(n, coords)
	if res != 3 {
		t.Errorf("the sample should give 3, but got %d", res)
	}
}
