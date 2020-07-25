package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, C []int, expect int) {
	res := solve(n, E, C)

	if res != expect {
		t.Errorf("Sample %d %v %v, expect %d, but got %d", n, E, C, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	E := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{3, 5},
		{3, 6},
	}
	C := []int{0, 1, 0, 1, 2, 0}
	expect := 5
	runSample(t, n, E, C, expect)
}

func TestSample2(t *testing.T) {
	n := 6
	E := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{3, 5},
		{3, 6},
	}
	C := []int{1, 0, 0, 0, 2, 0}
	expect := 2
	runSample(t, n, E, C, expect)
}

func TestSample3(t *testing.T) {
	n := 3
	E := [][]int{
		{1, 2},
		{2, 3},
	}
	C := []int{1, 0, 1}
	expect := 1
	runSample(t, n, E, C, expect)
}
