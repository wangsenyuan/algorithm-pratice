package main

import "testing"

func runSample(t *testing.T, n, K int, A []int, E [][]int, expect int64) {
	res := solve(n, K, A, E)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 3, 1
	A := []int{0, 1, 2}
	E := [][]int{
		{1, 2},
		{1, 3},
	}
	var expect int64 = 1
	runSample(t, n, k, A, E, expect)
}

func TestSample2(t *testing.T) {
	n, k := 4, 0
	A := []int{0, 1, 2, 3}
	E := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
	}
	var expect int64 = 2
	runSample(t, n, k, A, E, expect)
}

func TestSample3(t *testing.T) {
	n, k := 7, 0
	A := []int{0, 2, 1, 4, 2, 7, 3}
	E := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
		{2, 5},
		{3, 6},
		{3, 7},
	}
	var expect int64 = 6
	runSample(t, n, k, A, E, expect)
}
