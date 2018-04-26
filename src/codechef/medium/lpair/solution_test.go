package main

import "testing"

func runSample(t *testing.T, n int, pairs [][]int, expect int64) {
	res := solve(n, pairs)
	if res != expect {
		t.Errorf("sample %d %v, expect %d, but got %d", n, pairs, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	pairs := [][]int{
		{2, 3},
		{3, 6},
		{5, 4},
	}
	runSample(t, n, pairs, 1)
}

func TestSample2(t *testing.T) {
	n := 4
	pairs := [][]int{
		{5, 12},
		{10, 11},
		{11, 9},
		{30, 1},
	}
	runSample(t, n, pairs, 6)
}

func TestSample3(t *testing.T) {
	n := 4
	pairs := [][]int{
		{4, 10},
		{1, 30},
		{2, 20},
		{3, 5},
	}
	runSample(t, n, pairs, 5)
}



func TestSample4(t *testing.T) {
	n := 4
	pairs := [][]int{
		{4, 10},
		{1, 1},
		{2, 2},
		{3, 5},
	}
	runSample(t, n, pairs, 0)
}
