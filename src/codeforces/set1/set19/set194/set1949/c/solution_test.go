package main

import "testing"

func runSample(t *testing.T, n int, branches [][]int, expect bool) {
	res := solve(n, branches)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 7
	branches := [][]int{
		{5, 1},
		{3, 2},
		{4, 6},
		{3, 6},
		{7, 1},
		{1, 3},
	}
	expect := true
	runSample(t, n, branches, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	branches := [][]int{
		{1, 4},
		{4, 2},
		{3, 2},
		{5, 3},
	}
	expect := false
	runSample(t, n, branches, expect)
}

func TestSample3(t *testing.T) {
	n := 6
	branches := [][]int{
		{4, 5},
		{5, 6},
		{6, 1},
		{2, 6},
		{3, 2},
	}
	expect := true
	runSample(t, n, branches, expect)
}
