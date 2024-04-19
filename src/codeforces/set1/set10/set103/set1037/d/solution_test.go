package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, ord []int, expect bool) {
	res := solve(n, edges, ord)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	edges := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
	}
	ord := []int{1, 2, 3, 4}
	runSample(t, n, edges, ord, true)
}

func TestSample2(t *testing.T) {
	n := 4
	edges := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
	}
	ord := []int{1, 2, 4, 3}
	runSample(t, n, edges, ord, false)
}
