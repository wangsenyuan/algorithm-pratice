package main

import "testing"

func runSample(t *testing.T, n int, P [][]int, expect int64) {
	res := solve(n, P)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 1
	P := [][]int{
		{100, 100},
	}
	var expect int64
	runSample(t, n, P, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	P := [][]int{
		{1, 100},
		{100, 100},
		{100, 1},
		{1, 1},
	}
	var expect int64
	runSample(t, n, P, expect)
}

func TestSample3(t *testing.T) {
	n := 5
	P := [][]int{
		{1, 100},
		{100, 100},
		{100, 1},
		{1, 1},
		{50, 50},
	}
	var expect int64 = 4851
	runSample(t, n, P, expect)
}

func TestSample4(t *testing.T) {
	n := 19
	P := [][]int{
		{1, 1},
		{2, 1},
		{3, 1},
		{1, 2},
		{1, 3},
		{1, 4},
		{1, 5},
		{3, 2},
		{3, 3},
		{3, 4},
		{3, 5},
		{3, 6},
		{3, 7},
		{3, 8},
		{4, 8},
		{5, 8},
		{5, 7},
		{5, 6},
		{5, 5},
	}
	var expect int64 = 14
	runSample(t, n, P, expect)
}

func TestSample5(t *testing.T) {
	n := 20
	P := [][]int{
		{3, 5},
		{9, 0},
		{5, 2},
		{3, 1},
		{9, 6},
		{0, 6},
		{5, 6},
		{7, 7},
		{2, 9},
		{4, 7},
		{6, 3},
		{2, 3},
		{7, 6},
		{8, 6},
		{6, 4},
		{1, 7},
		{2, 5},
		{9, 4},
		{5, 5},
		{2, 2},
	}
	var expect int64 = 56
	runSample(t, n, P, expect)
}
