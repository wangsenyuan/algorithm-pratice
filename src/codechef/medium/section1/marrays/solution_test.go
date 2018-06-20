package main

import "testing"

func runSample(t *testing.T, n int, dishs [][]int, expect int64) {
	res := solve(n, dishs)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, dishs, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	dishs := [][]int{
		{1, 2, 3},
		{3, 2},
		{4, 5},
	}
	runSample(t, n, dishs, 8)
}

func TestSample2(t *testing.T) {
	n := 2
	dishs := [][]int{
		{1, 2},
		{4, 5},
	}
	runSample(t, n, dishs, 4)
}

func TestSample3(t *testing.T) {
	n := 3
	dishs := [][]int{
		{1, 2},
		{3, 2, 5},
		{4, 5},
	}
	runSample(t, n, dishs, 10)
}
