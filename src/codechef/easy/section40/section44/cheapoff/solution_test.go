package main

import "testing"

func runSample(t *testing.T, A []int, k int, E [][]int, expect int) {
	res := solve(A, k, E)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{5, 6, 7}
	k := 2
	E := [][]int{
		{1, 2},
		{2, 3},
	}
	expect := 1
	runSample(t, A, k, E, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 2, 4, 8}
	k := 3
	E := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 1},
	}
	expect := 7
	runSample(t, A, k, E, expect)
}

func TestSample3(t *testing.T) {
	A := []int{221, 373, 34, 344, 2344, 9848, 6666, 2323}
	k := 5
	E := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{2, 3},
		{2, 5},
		{2, 6},
		{7, 8},
		{6, 7},
		{5, 6},
		{4, 5},
	}
	expect := 2559
	runSample(t, A, k, E, expect)
}

func TestSample4(t *testing.T) {
	A := []int{5, 6, 7}
	k := 1
	E := [][]int{
		{1, 2},
		{2, 3},
	}
	expect := 0
	runSample(t, A, k, E, expect)
}
