package main

import "testing"

func runSample(t *testing.T, n, m int, queens [][]int, expect int) {
	res := solve(n, m, queens)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 5, 2
	queens := [][]int{
		{1, 1},
		{3, 1},
	}
	expect := 1
	runSample(t, n, m, queens, expect)
}
