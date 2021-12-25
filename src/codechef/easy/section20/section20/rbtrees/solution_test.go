package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, C string, expect int) {
	res := solve(n, E, C)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 7
	E := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
		{2, 5},
		{3, 6},
		{3, 7},
	}
	C := "0010010"
	expect := 3
	runSample(t, n, E, C, expect)
}
