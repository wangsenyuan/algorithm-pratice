package main

import "testing"

func runSample(t *testing.T, n int, H []int, E [][]int, expect int64) {
	res := solve(n, H, E)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	H := []int{1, 2, 1}
	E := [][]int{
		{1, 2},
		{2, 3},
	}
	var expect int64 = 4
	runSample(t, n, H, E, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	H := []int{1, 3, 3, 1, 3}
	E := [][]int{
		{1, 3},
		{5, 4},
		{4, 3},
		{2, 3},
	}
	var expect int64 = 7
	runSample(t, n, H, E, expect)
}

func TestSample3(t *testing.T) {
	n := 2
	H := []int{6, 1}
	E := [][]int{
		{1, 2},
	}
	var expect int64 = 12
	runSample(t, n, H, E, expect)
}

func TestSample4(t *testing.T) {
	n := 6
	H := []int{8, 8, 11, 3, 8, 9}
	E := [][]int{
		{6, 5},
		{2, 5},
		{1, 3},
		{3, 6},
		{5, 4},
	}
	var expect int64 = 25
	runSample(t, n, H, E, expect)
}

func TestSample5(t *testing.T) {
	n := 8
	H := []int{1, 1, 16, 16, 1, 1, 7, 7}
	E := [][]int{
		{1, 3},
		{4, 6},
		{4, 8},
		{2, 4},
		{2, 1},
		{7, 4},
		{5, 7},
	}
	var expect int64 = 40
	runSample(t, n, H, E, expect)
}

func TestSample6(t *testing.T) {
	n := 10
	H := []int{4, 8, 9, 4, 9, 9, 1, 6, 10, 9}
	E := [][]int{
		{8, 7},
		{8, 6},
		{5, 9},
		{3, 6},
		{9, 2},
		{4, 9},
		{7, 2},
		{1, 8},
		{10, 8},
	}
	var expect int64 = 37
	runSample(t, n, H, E, expect)
}
