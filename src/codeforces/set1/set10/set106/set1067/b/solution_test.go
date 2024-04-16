package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, k int, expect bool) {
	res := solve(n, edges, k)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 14
	edges := [][]int{
		{1, 4},
		{2, 4},
		{3, 4},
		{4, 13},
		{10, 5},
		{11, 5},
		{12, 5},
		{14, 5},
		{5, 13},
		{6, 7},
		{8, 6},
		{13, 6},
		{9, 6},
	}
	k := 2
	expect := true
	runSample(t, n, edges, k, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	edges := [][]int{
		{1, 3},
		{2, 3},
	}
	k := 1
	expect := false
	runSample(t, n, edges, k, expect)
}

func TestSample3(t *testing.T) {
	n := 8
	edges := [][]int{
		{8, 2},
		{2, 5},
		{5, 1},
		{7, 2},
		{2, 4},
		{3, 5},
		{5, 6},
	}
	k := 1
	expect := false
	runSample(t, n, edges, k, expect)
}

func TestSample4(t *testing.T) {
	n := 2
	edges := [][]int{
		{1, 2},
	}
	k := 1
	expect := false
	runSample(t, n, edges, k, expect)
}
