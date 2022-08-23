package main

import "testing"

func runSample(t *testing.T, C int, E [][]int, expect int) {
	res := solve(C, E)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	C := 3
	E := [][]int{
		{1, 2},
		{1, 3},
	}
	expect := 6
	runSample(t, C, E, expect)
}
