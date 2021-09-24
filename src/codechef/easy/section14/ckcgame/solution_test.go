package main

import "testing"

func runSample(t *testing.T, n int, V [][]int, A, B []int, expect int64) {
	res := solve(n, V, A, B)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	A := []int{1, 1}
	B := []int{2, 2}
	V := [][]int{
		{1, 10},
		{5, 6},
	}
	var expect int64 = 10
	runSample(t, n, V, A, B, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	A := []int{4, 4}
	B := []int{1, 1}
	V := [][]int{
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 10, 1},
		{1, 1, 1, 10},
	}
	var expect int64 = 30
	runSample(t, n, V, A, B, expect)
}

func TestSample3(t *testing.T) {
	n := 4
	A := []int{1, 3}
	B := []int{1, 4}
	V := [][]int{
		{1, 2, 3, 4},
		{2, 1, 4, 3},
		{3, 4, 1, 2},
		{4, 3, 2, 1},
	}
	var expect int64 = 8
	runSample(t, n, V, A, B, expect)
}
