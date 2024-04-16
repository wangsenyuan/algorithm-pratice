package main

import "testing"

func runSample(t *testing.T, n, k int, tariffs [][]int, expect int) {
	res := solve(n, tariffs, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 5, 7
	tariffs := [][]int{
		{1, 4, 5, 3},
		{1, 3, 5, 2},
		{2, 5, 10, 1},
	}
	expect := 44
	runSample(t, n, k, tariffs, expect)
}

func TestSample2(t *testing.T) {
	n, k := 7, 13
	tariffs := [][]int{
		{2, 3, 10, 7},
		{3, 5, 10, 10},
		{1, 2, 10, 6},
		{4, 5, 10, 9},
		{3, 4, 10, 8},
	}
	expect := 462
	runSample(t, n, k, tariffs, expect)
}

func TestSample3(t *testing.T) {
	n, k := 4, 100
	tariffs := [][]int{
		{3, 3, 2, 5},
		{1, 1, 3, 2},
		{2, 4, 4, 4},
	}
	expect := 64
	runSample(t, n, k, tariffs, expect)
}
