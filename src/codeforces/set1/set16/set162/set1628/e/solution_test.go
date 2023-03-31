package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, E [][]int, Q [][]int, expect []int) {
	res := solve(n, E, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	E := [][]int{
		{1, 3, 1},
		{2, 3, 2},
		{4, 5, 3},
		{4, 6, 4},
		{3, 4, 5},
	}
	Q := [][]int{
		{3, 1},
		{1, 1, 1},
		{3, 1},
		{2, 1, 1},
		{1, 5, 6},
		{3, 4},
		{2, 6, 6},
		{3, 4},
		{3, 1},
	}
	expect := []int{-1, -1, 4, 3, 5}
	runSample(t, n, E, Q, expect)
}
