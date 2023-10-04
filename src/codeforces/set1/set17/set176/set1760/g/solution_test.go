package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, a, b int, expect bool) {
	res := solve(n, E, a, b)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a, b := 1, 4
	n := 5
	E := [][]int{
		{1, 3, 1},
		{2, 3, 2},
		{4, 3, 3},
		{3, 5, 1},
	}
	expect := true
	runSample(t, n, E, a, b, expect)
}

func TestSample2(t *testing.T) {
	a, b := 1, 2
	n := 2
	E := [][]int{
		{1, 2, 2},
	}
	expect := false
	runSample(t, n, E, a, b, expect)
}

func TestSample3(t *testing.T) {
	a, b := 2, 3
	n := 6
	E := [][]int{
		{1, 2, 1},
		{2, 3, 1},
		{3, 4, 1},
		{4, 5, 3},
		{5, 6, 5},
	}
	expect := true
	runSample(t, n, E, a, b, expect)
}
