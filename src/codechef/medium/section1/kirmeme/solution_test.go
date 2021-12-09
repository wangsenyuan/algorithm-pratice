package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, A []int, L, R int, expect int) {
	res := solve(n, E, A, L, R)

	if int(res) != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	L, R := 1, 2
	A := []int{1, 3, 1}
	E := [][]int{
		{1, 2},
		{2, 3},
	}
	expect := 1
	runSample(t, n, E, A, L, R, expect)
}

func TestSample2(t *testing.T) {
	n := 7
	L, R := 2, 2
	A := []int{1, 8, 2, 3, 4, 8, 1}
	E := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{3, 5},
		{3, 6},
		{6, 7},
	}
	expect := 1
	runSample(t, n, E, A, L, R, expect)
}

func TestSample3(t *testing.T) {
	n := 9
	L, R := 1, 7
	A := []int{5, 1, 1, 1, 1, 1, 1, 1, 1}
	E := [][]int{
		{1, 2},
		{2, 3},
		{1, 4},
		{4, 5},
		{1, 6},
		{6, 7},
		{1, 8},
		{8, 9},
	}
	expect := 24
	runSample(t, n, E, A, L, R, expect)
}
