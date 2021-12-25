package main

import "testing"

func runSample(t *testing.T, n, m int, A [][]int, expect int) {
	res := solve(n, m, A)
	if res != expect {
		t.Errorf("sample %d %d %v, expect %d, but got %d", n, m, A, expect, res)
	}
}

func TestSample(t *testing.T) {
	n, m := 4, 4
	A := [][]int{
		{1, 2, 1, 2},
		{3, 4, 1, 2},
		{1, 1, 1, 1},
		{2, 2, 2, 2},
	}

	expect := 6
	runSample(t, n, m, A, expect)
}
