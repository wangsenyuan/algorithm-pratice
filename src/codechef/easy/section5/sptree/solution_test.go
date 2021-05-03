package main

import "testing"

func runSample(t *testing.T, n int, r int, E [][]int, F []int, expect []int) {
	res := solve(n, r, E, F)

	for i := 0; i < len(res); i++ {
		if res[i].first != expect[i] {
			t.Errorf("Sample result %v, not correct, as expected %v", res, expect)
		}
	}
}

func TestSample1(t *testing.T) {
	n, r := 5, 3
	E := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
		{2, 5},
	}

	F := []int{2}

	expect := []int{1, 2, 0, 1, 1}

	runSample(t, n, r, E, F, expect)
}

func TestSample2(t *testing.T) {
	n, r := 8, 2
	E := [][]int{
		{1, 2},
		{2, 3},
		{2, 4},
		{2, 5},
		{4, 6},
		{5, 7},
		{5, 8},
	}

	F := []int{6, 5, 8}

	expect := []int{-1, 0, -1, 1, 1, 2, 0, 2}

	runSample(t, n, r, E, F, expect)
}
