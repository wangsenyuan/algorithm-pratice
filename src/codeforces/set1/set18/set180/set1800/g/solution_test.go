package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, expect bool) {
	res := solve(n, E)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	E := [][]int{
		{1, 5},
		{1, 6},
		{1, 2},
		{2, 3},
		{2, 4},
	}
	expect := true
	runSample(t, n, E, expect)
}

func TestSample2(t *testing.T) {
	n := 7
	E := [][]int{
		{1, 5},
		{1, 3},
		{3, 6},
		{1, 4},
		{4, 7},
		{4, 2},
	}
	expect := false
	runSample(t, n, E, expect)
}

func TestSample3(t *testing.T) {
	n := 9
	E := [][]int{
		{1, 2},
		{2, 4},
		{2, 3},
		{3, 5},
		{1, 7},
		{7, 6},
		{7, 8},
		{8, 9},
	}
	expect := true
	runSample(t, n, E, expect)
}
