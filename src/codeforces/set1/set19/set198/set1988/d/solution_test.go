package main

import "testing"

func runSample(t *testing.T, a []int, edges [][]int, expect int) {
	res := solve(a, edges)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1000000000000}
	expect := 1000000000000
	runSample(t, a, nil, expect)
}

func TestSample2(t *testing.T) {
	a := []int{47, 15, 32, 29, 23}
	edges := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
		{2, 5},
	}
	expect := 193
	runSample(t, a, edges, expect)
}

func TestSample3(t *testing.T) {
	a := []int{8, 10, 2, 3, 5, 7, 4}
	edges := [][]int{
		{1, 2},
		{1, 4},
		{3, 2},
		{5, 3},
		{6, 2},
		{7, 5},
	}
	expect := 57
	runSample(t, a, edges, expect)
}
