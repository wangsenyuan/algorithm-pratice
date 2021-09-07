package main

import "testing"

func runSample(t *testing.T, n int, k int, S []string, A [][]int, expect int64) {
	res := solve(n, k, S, A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	k := 1
	S := []string{
		".#",
		"#.",
	}
	A := [][]int{
		{1, 2},
		{3, 4},
	}
	var expect int64 = 8
	runSample(t, n, k, S, A, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	k := 3
	S := []string{
		".#.",
		"#.#",
		"..#",
	}
	A := [][]int{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	}
	var expect int64 = -1
	runSample(t, n, k, S, A, expect)
}

func TestSample3(t *testing.T) {
	n := 5
	k := 4
	S := []string{
		".....",
		"#####",
		".....",
		"#####",
		".....",
	}
	A := [][]int{
		{2, 3, 4, 1, 5},
		{1, 7, 15, 12, 2},
		{2, 5, 10, 8, 3},
		{9, 9, 9, 9, 9},
		{1, 2, 3, 4, 5},
	}
	var expect int64 = 62
	runSample(t, n, k, S, A, expect)
}
