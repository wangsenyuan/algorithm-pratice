package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, x int, y int, A []int, expect int) {
	res := solve(n, E, x, y, A)

	if int(res) != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	x, y := 1, 3
	n := 3
	E := [][]int{
		{1, 3},
		{1, 2},
	}
	A := []int{2}
	expect := 3
	runSample(t, n, E, x, y, A, expect)
}

func TestSample2(t *testing.T) {
	x, y := 3, 5
	n := 6
	E := [][]int{
		{1, 3},
		{3, 4},
		{3, 5},
		{5, 6},
		{5, 2},
	}
	A := []int{1, 6, 2, 1}
	expect := 7
	runSample(t, n, E, x, y, A, expect)
}

func TestSample3(t *testing.T) {
	x, y := 3, 2
	n := 6
	E := [][]int{
		{1, 3},
		{3, 4},
		{3, 5},
		{5, 6},
		{5, 2},
	}
	A := []int{5, 3}
	expect := 2
	runSample(t, n, E, x, y, A, expect)
}
