package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, expect bool) {
	res := solve(n, E)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	E := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{1, 5},
	}
	expect := false
	runSample(t, n, E, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	E := [][]int{
		{1, 3},
		{1, 2},
		{4, 5},
		{3, 4},
	}
	expect := true
	runSample(t, n, E, expect)
}

func TestSample3(t *testing.T) {
	n := 8
	E := [][]int{
		{4, 6},
		{6, 7},
		{2, 5},
		{4, 5},
		{3, 4},
		{1, 2},
		{7, 8},
	}
	expect := true
	runSample(t, n, E, expect)
}
