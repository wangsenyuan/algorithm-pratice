package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, h []int, p []int, expect int) {
	res := solve(n, E, h, p)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	E := [][]int{
		{1, 2},
		{2, 3},
		{2, 4},
		{3, 5},
		{4, 5},
		{3, 6},
		{6, 5},
	}
	h := []int{2, 3, 4, 5, 6}
	p := []int{1, 2, 3, 5}
	expect := 2
	runSample(t, n, E, h, p, expect)
}

func TestSample2(t *testing.T) {
	n := 6
	E := [][]int{
		{1, 2},
		{2, 3},
		{2, 4},
		{3, 5},
		{4, 5},
		{3, 6},
		{6, 5},
	}
	h := []int{2, 3, 4, 5, 6, 5}
	p := []int{1, 2, 3, 5}
	expect := 1
	runSample(t, n, E, h, p, expect)
}

func TestSample3(t *testing.T) {
	n := 4
	E := [][]int{
		{1, 2},
		{1, 3},
		{2, 3},
		{3, 4},
	}
	h := []int{3, 4, 2}
	p := []int{1, 3}
	expect := 1
	runSample(t, n, E, h, p, expect)
}
