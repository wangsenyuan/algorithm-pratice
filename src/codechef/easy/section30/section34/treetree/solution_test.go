package main

import "testing"

func runSample(t *testing.T, n int, T [][][]int, A []int, expect int) {
	res := solve(n, T, A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	T := [][][]int{
		{
			{1, 2, 1},
			{2, 3, 3},
		},
		{
			{1, 2, 2},
			{1, 3, 1},
			{1, 4, 3},
		},
	}
	A := []int{1}
	expect := 72
	runSample(t, n, T, A, expect)
}
