package main

import "testing"

func runSample(t *testing.T, n int, k []int, edges [][]int, r int, expect int) {
	res := solve(n, k, edges, r)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	k := []int{1, 3, 1, 3, 2}
	edges := [][]int{
		{2, 5},
		{3, 4},
		{4, 5},
		{1, 5},
	}
	s := 4
	expect := 6
	runSample(t, n, k, edges, s, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	k := []int{2, 1, 1}
	edges := [][]int{
		{3, 2},
		{1, 2},
	}
	s := 3
	expect := 2
	runSample(t, n, k, edges, s, expect)
}

func TestSample3(t *testing.T) {
	n := 1
	k := []int{6}

	s := 1
	expect := 0
	runSample(t, n, k, nil, s, expect)
}
