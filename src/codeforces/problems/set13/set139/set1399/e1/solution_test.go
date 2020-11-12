package main

import "testing"

func runSample(t *testing.T, n int, S int64, E [][]int, expect int) {
	res := solve(n, S, E)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	var S int64 = 20
	E := [][]int{
		{2, 1, 8},
		{3, 1, 7},
	}
	expect := 0
	runSample(t, n, S, E, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	var S int64 = 50
	E := [][]int{
		{1, 3, 100},
		{1, 5, 10},
		{2, 3, 123},
		{5, 4, 55},
	}
	expect := 8
	runSample(t, n, S, E, expect)
}

func TestSample3(t *testing.T) {
	n := 2
	var S int64 = 100
	E := [][]int{
		{1, 2, 409},
	}
	expect := 3
	runSample(t, n, S, E, expect)
}

func TestSample4(t *testing.T) {
	n := 10
	var S int64 = 28
	E := [][]int{
		{8, 2, 8},
		{5, 1, 4},
		{6, 1, 10},
		{10, 2, 7},
		{7, 2, 1},
		{9, 2, 1},
		{2, 1, 5},
		{4, 1, 9},
		{3, 2, 5},
	}
	expect := 7
	runSample(t, n, S, E, expect)
}
