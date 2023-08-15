package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, E [][]int, Q [][]int, expect []int64) {
	res := solve(n, E, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 8
	E := [][]int{
		{6, 7},
		{3, 2},
		{8, 3},
		{5, 7},
		{7, 4},
		{7, 1},
		{7, 3},
	}
	Q := [][]int{
		{1, 0},
		{1, 2},
		{1, 3},
		{7, 1},
		{5, 0},
		{7, 200000},
	}
	expect := []int64{5, 2, 1, 4, 0, 4}
	runSample(t, n, E, Q, expect)
}
