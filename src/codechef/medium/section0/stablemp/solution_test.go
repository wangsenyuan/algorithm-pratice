package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, W [][]int, M [][]int, expect []int) {
	res := solve(n, W, M)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v %v, expect %v, but got %v", n, W, M, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	W := [][]int{
		{1, 4, 3, 1, 2},
		{2, 2, 1, 3, 4},
		{3, 1, 3, 4, 2},
		{4, 4, 3, 1, 2},
	}

	M := [][]int{
		{1, 3, 2, 4, 1},
		{2, 2, 3, 1, 4},
		{3, 3, 1, 2, 4},
		{4, 3, 2, 4, 1},
	}

	expect := []int{3, 2, 1, 4}
	runSample(t, n, W, M, expect)
}

func TestSample2(t *testing.T) {
	n := 7
	W := [][]int{
		{1, 3, 4, 2, 1, 6, 7, 5},
		{2, 6, 4, 2, 3, 5, 1, 7},
		{3, 6, 3, 5, 7, 2, 4, 1},
		{4, 1, 6, 3, 2, 4, 7, 5},
		{5, 1, 6, 5, 3, 4, 7, 2},
		{6, 1, 7, 3, 4, 5, 6, 2},
		{7, 5, 6, 2, 4, 3, 7, 1},
	}

	M := [][]int{
		{1, 4, 5, 3, 7, 2, 6, 1},
		{2, 5, 6, 4, 7, 3, 2, 1},
		{3, 1, 6, 5, 4, 3, 7, 2},
		{4, 3, 5, 6, 7, 2, 4, 1},
		{5, 1, 7, 6, 4, 3, 5, 2},
		{6, 6, 3, 7, 5, 2, 4, 1},
		{7, 1, 7, 4, 2, 6, 5, 3},
	}

	expect := []int{4, 5, 1, 3, 7, 6, 2}
	runSample(t, n, W, M, expect)
}
