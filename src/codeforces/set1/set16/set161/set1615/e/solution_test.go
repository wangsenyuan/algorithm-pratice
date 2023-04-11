package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, k int, expect int64) {
	res := solve(n, E, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 4, 2
	E := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
	}
	var expect int64 = 1
	runSample(t, n, E, k, expect)
}

func TestSample2(t *testing.T) {
	n, k := 5, 2
	E := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
	}
	var expect int64 = 6
	runSample(t, n, E, k, expect)
}

func TestSample3(t *testing.T) {
	n, k := 7, 2
	E := [][]int{
		{1, 2},
		{1, 3},
		{4, 2},
		{3, 5},
		{6, 3},
		{6, 7},
	}
	var expect int64 = 4
	runSample(t, n, E, k, expect)
}

func TestSample4(t *testing.T) {
	n, k := 4, 1
	E := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
	}
	var expect int64 = -1
	runSample(t, n, E, k, expect)
}
