package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, expect bool) {
	res := solve(n, edges)
	if res != expect {
		t.Errorf("sample %d %v, expect %t, but got %t", n, edges, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	edges := [][]int{
		{1, 2},
		{2, 3},
	}
	expect := true
	runSample(t, n, edges, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	edges := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{3, 5},
	}
	expect := true
	runSample(t, n, edges, expect)
}
