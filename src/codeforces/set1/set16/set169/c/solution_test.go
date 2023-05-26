package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, expect int) {
	res := solve(n, E)

	if res != expect {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	E := [][]int{{1, 2}}
	expect := 1
	runSample(t, n, E, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	E := [][]int{
		{1, 2},
		{1, 4},
		{2, 4},
		{1, 4},
	}
	expect := 2
	runSample(t, n, E, expect)
}

func TestSample3(t *testing.T) {
	n := 5
	E := [][]int{
		{1, 2},
		{2, 3},
		{3, 5},
		{1, 4},
		{4, 3},
		{4, 5},
		{3, 1},
	}
	expect := 4
	runSample(t, n, E, expect)
}
