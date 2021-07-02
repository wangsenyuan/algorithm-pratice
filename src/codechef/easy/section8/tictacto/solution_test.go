package main

import "testing"

func runSample(t *testing.T, m, n, k int, cmds [][]int, expect string) {
	res := solve(m, n, k, cmds)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	m, n, k := 3, 4, 2
	cmds := [][]int{
		{1, 1},
		{3, 1},
		{1, 2},
		{2, 3},
		{2, 1},
		{2, 4},
		{3, 2},
		{3, 4},
		{1, 3},
		{3, 3},
		{2, 2},
		{1, 4},
	}
	expect := Bob

	runSample(t, m, n, k, cmds, expect)
}
