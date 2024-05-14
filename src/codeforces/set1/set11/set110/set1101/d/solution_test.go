package main

import "testing"

func runSample(t *testing.T, a []int, edges [][]int, expect int) {
	res := solve(a, edges)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{2, 3, 4}
	edges := [][]int{
		{1, 2},
		{2, 3},
	}
	expect := 1
	runSample(t, a, edges, expect)
}

func TestSample2(t *testing.T) {
	a := []int{2, 3, 4}
	edges := [][]int{
		{1, 3},
		{2, 3},
	}
	expect := 2
	runSample(t, a, edges, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 1, 1}
	edges := [][]int{
		{1, 3},
		{2, 3},
	}
	expect := 0
	runSample(t, a, edges, expect)
}
