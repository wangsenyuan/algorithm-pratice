package main

import "testing"

func runSample(t *testing.T, cats [][]int, n int, expect int) {
	res := solve(cats, n)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 15
	cats := [][]int{
		{2, 10},
		{3, 5},
		{2, 4},
		{7, 7},
		{8, 12},
		{11, 11},
	}
	expect := 5
	runSample(t, cats, n, expect)
}

func TestSample2(t *testing.T) {
	n := 1000
	cats := [][]int{
		{1, 1000},
	}
	expect := 1
	runSample(t, cats, n, expect)
}

func TestSample3(t *testing.T) {
	n := 5
	cats := [][]int{
		{1, 2},
		{3, 4},
		{3, 4},
		{3, 4},
		{3, 4},
		{1, 1},
		{1, 2},
		{3, 3},
		{3, 4},
		{3, 4},
	}
	expect := 10
	runSample(t, cats, n, expect)
}
