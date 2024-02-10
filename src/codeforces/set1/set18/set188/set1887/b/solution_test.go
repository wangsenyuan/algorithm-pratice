package main

import "testing"

func runSample(t *testing.T, n int, moments [][][]int, a []int, expect int) {
	res := solve(n, moments, a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	moments := [][][]int{
		{
			{1, 2},
			{2, 3},
			{3, 4},
			{4, 5},
		},
		{
			{2, 3},
			{3, 5},
		},
	}
	a := []int{2, 1, 2, 1, 2, 1}
	expect := 5
	runSample(t, n, moments, a, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	moments := [][][]int{
		{
			{1, 2},
			{3, 1},
			{4, 3},
		},
		{
			{2, 1},
			{4, 5},
		},
	}
	a := []int{1, 2, 1, 1, 1}
	expect := -1
	runSample(t, n, moments, a, expect)
}
