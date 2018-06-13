package main

import "testing"

func runSample(t *testing.T, n, m int, edges [][]int, expect int) {
	res := solve(n, m, edges)

	if res != expect {
		t.Errorf("sample %d %d %v, expect %d, but got %d", n, m, edges, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 3, 3
	edges := [][]int{
		{1, 2},
		{2, 3},
		{1, 3},
	}
	expect := 1
	runSample(t, n, m, edges, expect)
}

func TestSample2(t *testing.T) {
	n, m := 4, 3
	edges := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
	}
	expect := 2
	runSample(t, n, m, edges, expect)
}
