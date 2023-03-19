package main

import "testing"

func runSample(t *testing.T, n int, m int, ladders [][]int, X []int, ok bool, expect int64) {
	ok1, res := solve(n, m, ladders, X)

	if ok != ok1 || ok && expect != res {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 5, 3
	X := []int{5, 17, 8, 1, 4}
	ladders := [][]int{
		{1, 3, 3, 3, 4},
		{3, 1, 5, 2, 5},
		{3, 2, 5, 1, 6},
	}
	ok := true
	var expect int64 = 16
	runSample(t, n, m, ladders, X, ok, expect)
}

func TestSample2(t *testing.T) {
	n, m := 6, 3
	X := []int{5, 17, 8, 1, 4, 2}
	ladders := [][]int{
		{1, 3, 3, 3, 4},
		{3, 1, 5, 2, 5},
		{3, 2, 5, 1, 6},
	}
	ok := false
	var expect int64 = 16
	runSample(t, n, m, ladders, X, ok, expect)
}

func TestSample3(t *testing.T) {
	n, m := 5, 3
	X := []int{5, 17, 8, 1, 4}
	ladders := [][]int{
		{1, 3, 5, 3, 100},
	}
	ok := true
	var expect int64 = -90
	runSample(t, n, m, ladders, X, ok, expect)
}

func TestSample4(t *testing.T) {
	n, m := 5, 5
	X := []int{3, 2, 3, 7, 5}
	ladders := [][]int{
		{3, 5, 4, 2, 1},
		{2, 2, 5, 4, 5},
		{4, 4, 5, 2, 3},
		{1, 2, 4, 2, 2},
		{3, 3, 5, 2, 4},
	}
	ok := true
	var expect int64 = 27
	runSample(t, n, m, ladders, X, ok, expect)
}
