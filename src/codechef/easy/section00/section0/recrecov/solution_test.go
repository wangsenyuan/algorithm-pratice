package main

import "testing"

func runSample(t *testing.T, n int, m int, pairs [][]int, expect int) {
	res := solve(n, m, pairs)
	if res != expect {
		t.Errorf("sample: %d %d %v, should give answer %d, but got %d", n, m, pairs, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 3, 3
	pairs := [][]int{
		{1, 2},
		{2, 3},
		{1, 3},
	}
	expect := 1
	runSample(t, n, m, pairs, expect)
}

func TestSample2(t *testing.T) {
	n, m := 3, 2
	pairs := [][]int{
		{1, 2},
		{1, 3},
	}
	expect := 2
	runSample(t, n, m, pairs, expect)
}

func TestSample3(t *testing.T) {
	n, m := 5, 4
	pairs := [][]int{
		{1, 3},
		{2, 3},
		{3, 4},
		{2, 5},
	}
	expect := 2
	runSample(t, n, m, pairs, expect)
}
