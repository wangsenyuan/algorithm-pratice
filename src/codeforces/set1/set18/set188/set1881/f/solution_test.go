package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, a []int, expect int) {
	res := solve(n, E, a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 7
	a := []int{2, 6, 7}
	E := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
		{2, 5},
		{3, 6},
		{3, 7},
	}
	expect := 2
	runSample(t, n, E, a, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	a := []int{1, 2, 3, 4}
	E := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
	}
	expect := 2
	runSample(t, n, E, a, expect)
}

func TestSample3(t *testing.T) {
	n := 10
	a := []int{1, 2, 4, 5, 6, 7, 8, 9, 10}
	E := [][]int{
		{1, 3},
		{3, 9},
		{9, 4},
		{4, 10},
		{10, 6},
		{6, 7},
		{7, 2},
		{2, 5},
		{5, 8},
	}
	expect := 5
	runSample(t, n, E, a, expect)
}
