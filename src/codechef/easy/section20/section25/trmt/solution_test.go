package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, V []int, E [][]int, Q [][]int, expect []int64) {
	res := solve(n, E, V, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	V := []int{5, 3, 3, 1, 2}
	E := [][]int{
		{1, 2, 2},
		{2, 3, 1},
		{2, 4, 1},
		{1, 5, 2},
	}
	Q := [][]int{
		{4, 5},
		{5, 4},
	}
	expect := []int64{1, -1}
	runSample(t, n, V, E, Q, expect)
}

func TestSample2(t *testing.T) {
	n := 10
	V := []int{3, 6, 9, 2, 3, 8, 3, 7, 9, 7}
	E := [][]int{
		{1, 2, 4},
		{1, 3, 6},
		{1, 4, 8},
		{4, 5, 9},
		{1, 6, 2},
		{5, 7, 3},
		{3, 8, 7},
		{3, 9, 2},
		{7, 10, 9},
	}
	Q := [][]int{
		{9, 2},
		{2, 1},
		{8, 5},
		{3, 2},
		{3, 10},
	}
	expect := []int64{-96, -12, -252, -24, -69}
	runSample(t, n, V, E, Q, expect)
}
