package main

import "testing"

func runSample(t *testing.T, N, M int, workers [][]int, expect int) {
	res := solve(N, M, workers)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 10, 3
	workers := [][]int{
		{1, 10},
		{1, 5},
		{6, 10},
	}
	expect := 3
	runSample(t, n, m, workers, expect)
}

func TestSample2(t *testing.T) {
	n, m := 10, 1
	workers := [][]int{
		{2, 10},
	}
	expect := -1
	runSample(t, n, m, workers, expect)
}

func TestSample3(t *testing.T) {
	n, m := 10, 2
	workers := [][]int{
		{5, 10},
		{1, 5},
	}
	expect := 5
	runSample(t, n, m, workers, expect)
}
