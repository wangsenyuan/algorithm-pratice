package main

import "testing"

func runSample(t *testing.T, N, H int, T []int, children [][]int, expect int) {
	res := solve(N, H, T, children)

	if res != expect {
		t.Errorf("Sample %d %d %v %v, expect %d, but got %d", N, H, T, children, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, H := 5, 24
	T := []int{13, 24, 22, 12, 16}
	children := [][]int{
		{1, 3},
		{0},
		{2, 2, 5},
		{0},
		{1, 4},
	}

	expect := 5
	runSample(t, N, H, T, children, expect)
}

func TestSample2(t *testing.T) {
	N, H := 10, 8
	T := []int{1, 4, 3, 1, 7, 3, 2, 2, 4, 4}
	children := [][]int{
		{3, 2, 5, 10},
		{2, 3, 4},
		{0},
		{0},
		{1, 6},
		{3, 7, 8, 9},
		{0},
		{0},
		{0},
		{0},
	}

	expect := 4
	runSample(t, N, H, T, children, expect)
}
