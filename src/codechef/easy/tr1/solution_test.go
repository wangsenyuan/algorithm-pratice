package main

import "testing"

func runSample(t *testing.T, roads [][]int, n int, expect int) {
	res := solve(n, roads)
	if res != expect {
		t.Errorf("sampel %v should give answer %d, but got %d", roads, expect, res)
	}
}

func TestSample1(t *testing.T) {
	roads := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
		{2, 5},
	}
	runSample(t, roads, 5, 41)
}

func TestSample2(t *testing.T) {
	roads := [][]int{
		{1, 2},
		{1, 3},
	}
	runSample(t, roads, 3, 12)
}

func TestSample3(t *testing.T) {
	roads := [][]int{}
	runSample(t, roads, 1, 1)
}
