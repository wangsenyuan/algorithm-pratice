package main

import "testing"

func runSample(t *testing.T, n int, people [][]int, expect int) {
	res := solve(n, people)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	people := [][]int{
		{1, 2},
		{4, 3},
		{1, 4},
		{3, 4},
	}
	expect := 1
	runSample(t, n, people, expect)
}

func TestSample2(t *testing.T) {
	n := 6
	people := [][]int{
		{2, 3},
		{2, 1},
		{3, 4},
		{6, 5},
		{4, 5},
	}
	expect := 0
	runSample(t, n, people, expect)
}

func TestSample3(t *testing.T) {
	n := 10
	people := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
		{5, 1},
		{1, 6},
		{2, 7},
		{3, 8},
		{4, 9},
		{5, 10},
		{6, 8},
		{7, 9},
		{8, 10},
		{9, 6},
		{10, 7},
	}
	expect := 6
	runSample(t, n, people, expect)
}
