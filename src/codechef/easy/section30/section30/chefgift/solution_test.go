package main

import "testing"

func runSample(t *testing.T, D int, B [][]int, C [][]int, expect int) {
	res := solve(D, B, C)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	D := 40
	B := [][]int{
		{10, 9, 8, 7, 6},
		{8, 7, 6, 9, 12},
		{16, 1, 1, 1, 1},
		{7, 6, 5, 4, 3},
	}
	C := [][]int{
		{0, 6, 17, 7},
		{8, 0, 19, 2},
		{9, 9, 0, 0},
		{12, 4, 18, 0},
	}
	expect := 20
	runSample(t, D, B, C, expect)
}

func TestSample2(t *testing.T) {
	D := 10
	B := [][]int{
		{5, 6},
		{6, 7},
	}
	C := [][]int{
		{0, 0},
		{0, 0},
	}
	expect := -1
	runSample(t, D, B, C, expect)
}

func TestSample3(t *testing.T) {
	D := 12
	B := [][]int{
		{3, 4, 5},
		{5, 7, 2},
	}
	C := [][]int{
		{0, 5},
		{2, 0},
	}
	expect := 0
	runSample(t, D, B, C, expect)
}
