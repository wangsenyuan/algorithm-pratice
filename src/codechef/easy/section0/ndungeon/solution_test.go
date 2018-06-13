package main

import "testing"

func TestSample(t *testing.T) {
	m, n := 4, 3
	rooms := [][]int{
		{2, 3, 2},
		{2, 5, 1},
		{5, 3, 1},
		{3, 1, 1},
	}
	a, b, x := 3, 1, 15

	spare, can := solve(m, n, rooms, a, b, x)
	if !can || spare != 4 {
		t.Errorf("the sample should give answer YES & 4, but got %v %d", can, spare)
	}
}
