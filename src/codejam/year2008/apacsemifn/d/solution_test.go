package main

import "testing"

func runSample(t *testing.T, N int, Large [][]int, M int, Small [][]int, expect bool) {
	res := solve(N, Large, M, Small)

	if res != expect {
		t.Errorf("Sample %d %v %d %v, expect %t, but got %t", N, Large, M, Small, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N := 5
	Large := [][]int{
		{1, 2}, {2, 3}, {3, 4}, {4, 5},
	}
	M := 4
	Small := [][]int{
		{1, 2}, {1, 3}, {1, 4},
	}
	expect := false
	runSample(t, N, Large, M, Small, expect)
}

func TestSample2(t *testing.T) {
	N := 5
	Large := [][]int{
		{1, 2}, {1, 3}, {1, 4}, {4, 5},
	}
	M := 4
	Small := [][]int{
		{1, 2}, {2, 3}, {3, 4},
	}
	expect := true
	runSample(t, N, Large, M, Small, expect)
}
