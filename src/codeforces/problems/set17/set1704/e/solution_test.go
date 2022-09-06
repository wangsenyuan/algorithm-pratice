package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, a []int, expect int) {
	res := solve(n, edges, a)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	a := []int{1, 1, 1}
	edges := [][]int{
		{1, 2},
		{2, 3},
	}
	expect := 3
	runSample(t, n, edges, a, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	a := []int{1, 0, 0, 0, 0}
	edges := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
		{1, 5},
	}
	expect := 5
	runSample(t, n, edges, a, expect)
}

func TestSample3(t *testing.T) {
	n := 10
	a := []int{998244353, 0, 0, 0, 998244353, 0, 0, 0, 0, 0}
	edges := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
		{5, 6},
		{6, 7},
		{7, 8},
		{8, 9},
		{9, 10},
		{1, 3},
		{7, 9},
	}
	expect := 4
	runSample(t, n, edges, a, expect)
}
