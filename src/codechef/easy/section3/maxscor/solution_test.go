package main

import "testing"

func runSample(t *testing.T, n, k int, subjects [][]int, expect int64) {
	res := solve(n, k, subjects)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 2, 2
	subs := [][]int{
		{1, 1},
		{2, 1},
	}
	var expect int64 = 2
	runSample(t, n, k, subs, expect)
}

func TestSample2(t *testing.T) {
	n, k := 3, 3
	subs := [][]int{
		{1, 1},
		{2, 2},
		{3, 3},
	}
	var expect int64 = 2
	runSample(t, n, k, subs, expect)
}

func TestSample3(t *testing.T) {
	n, k := 5, 6
	subs := [][]int{
		{5, 1},
		{4, 2},
		{3, 3},
		{2, 4},
		{1, 5},
	}
	var expect int64 = 9
	runSample(t, n, k, subs, expect)
}
