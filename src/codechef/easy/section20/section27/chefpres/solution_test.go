package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, N, K, B int, F []int, E [][]int, Q [][]int, expect []int) {
	res := solve(N, K, B, F, E, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k, b := 8, 4, 8
	E := [][]int{
		{3, 2},
		{2, 1},
		{1, 8},
		{8, 6},
		{8, 7},
		{7, 4},
		{4, 5},
	}
	F := []int{3, 2, 3, 4, 1, 1, 4, 4}

	Q := [][]int{
		{2, 1},
		{2, 2},
		{2, 3},
		{2, 4},
		{8, 1},
		{8, 2},
		{8, 3},
		{8, 4},
	}
	expect := []int{5, 2, 3, 4, 5, 2, 1, 4}

	runSample(t, n, k, b, F, E, Q, expect)
}

func TestSample2(t *testing.T) {
	n, k, b := 1, 2, 1
	E := [][]int{}
	F := []int{1}

	Q := [][]int{
		{1, 2},
	}
	expect := []int{-1}

	runSample(t, n, k, b, F, E, Q, expect)
}
