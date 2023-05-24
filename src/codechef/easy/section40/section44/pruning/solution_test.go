package main

import "testing"

func runSample(t *testing.T, d int, n int, E [][]int, expect int) {
	res := solve(d, n, E)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	d := 10
	E := [][]int{
		{0, 1, 10},
		{1, 2, 9},
		{1, 3, 10},
		{3, 4, 9},
	}
	expect := 29
	runSample(t, d, n, E, expect)
}

func TestSample2(t *testing.T) {
	n := 10
	d := 1
	E := [][]int{
		{0, 1, 1},
		{0, 2, 1},
		{1, 3, 1},
		{1, 4, 1},
		{2, 5, 1},
		{2, 6, 1},
		{2, 7, 1},
		{6, 8, 1},
		{6, 9, 1},
	}
	expect := 7
	runSample(t, d, n, E, expect)
}

func TestSample3(t *testing.T) {
	n := 7
	d := 2
	E := [][]int{
		{0, 1, 1},
		{0, 2, 2},
		{1, 3, 1},
		{1, 4, 2},
		{2, 5, 2},
		{2, 6, 7},
	}
	expect := 7
	runSample(t, d, n, E, expect)
}
