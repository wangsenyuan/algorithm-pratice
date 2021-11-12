package main

import "testing"

func runSample(t *testing.T, n int, T []int, E [][]int, expect int) {
	res := solve(n, T, E)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	T := []int{2, 2, -1, -2, 2}
	E := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{1, 5},
	}
	expect := 16
	runSample(t, n, T, E, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	T := []int{-1, -1, -1, -1, 2}
	E := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{1, 5},
	}
	expect := 2
	runSample(t, n, T, E, expect)
}
