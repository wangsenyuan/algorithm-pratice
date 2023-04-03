package main

import "testing"

func runSample(t *testing.T, A []int, E [][]int, expect int64) {
	res := solve(A, E)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{6, 3, 6, 7, 3, 3}
	E := [][]int{
		{3, 6},
	}
	var expect int64 = 40
	runSample(t, A, E, expect)
}

func TestSample2(t *testing.T) {
	A := []int{3, 4}
	E := [][]int{}
	var expect int64 = 14
	runSample(t, A, E, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 2, 2, 3, 1, 5, 1}
	E := [][]int{
		{1, 5},
		{3, 5},
		{1, 3},
		{2, 5},
	}
	var expect int64 = 15
	runSample(t, A, E, expect)
}
