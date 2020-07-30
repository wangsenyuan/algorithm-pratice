package main

import "testing"

func runSample(t *testing.T, n int, k int, A []int, E [][]int, expect int64) {
	res := solve(n, k, A, E)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 3, 3
	A := []int{1, 1, 1}
	E := [][]int{
		{1, 2},
		{2, 3},
	}
	var expect int64 = 3
	runSample(t, n, k, A, E, expect)
}

func TestSample2(t *testing.T) {
	n, k := 3, 3
	A := []int{1, 2, 3}
	E := [][]int{
		{1, 2},
		{2, 3},
	}
	var expect int64 = 8
	runSample(t, n, k, A, E, expect)
}
