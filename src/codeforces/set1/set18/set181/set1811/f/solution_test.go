package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, expect bool) {
	res := solve(n, E)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 16
	E := [][]int{
		{1, 2},
		{1, 5},
		{1, 6},
		{1, 7},
		{2, 3},
		{2, 7},
		{2, 8},
		{3, 10},
		{3, 11},
		{3, 4},
		{4, 13},
		{4, 14},
		{4, 15},
		{5, 6},
		{8, 9},
		{9, 10},
		{11, 12},
		{12, 13},
		{14, 16},
		{15, 16},
	}
	expect := false
	runSample(t, n, E, expect)
}

func TestSample2(t *testing.T) {
	n := 16
	E := [][]int{
		{1, 2},
		{2, 3},
		{1, 3},
		{1, 4},
		{1, 5},
		{5, 6},
		{6, 7},
		{7, 4},
		{4, 8},
		{8, 9},
		{9, 10},
		{10, 4},
		{2, 11},
		{11, 12},
		{12, 13},
		{13, 2},
		{3, 14},
		{14, 15},
		{15, 16},
		{16, 3},
	}
	expect := false
	runSample(t, n, E, expect)
}

func TestSample3(t *testing.T) {
	n := 9
	E := [][]int{
		{1, 2},
		{3, 1},
		{2, 3},
		{1, 6},
		{4, 1},
		{6, 4},
		{3, 8},
		{3, 5},
		{5, 8},
		{9, 7},
		{2, 9},
		{7, 2},
	}
	expect := true
	runSample(t, n, E, expect)
}
