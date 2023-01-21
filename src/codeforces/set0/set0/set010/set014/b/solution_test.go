package main

import "testing"

func runSample(t *testing.T, x0 int, men [][]int, expect int) {
	res := solve(x0, men)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	x0 := 3
	men := [][]int{
		{0, 7},
		{14, 2},
		{4, 6},
	}
	expect := 1
	runSample(t, x0, men, expect)
}
