package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, E [][]int, Q []int64, expect []int) {
	res := solve(n, E, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 8
	E := [][]int{
		{1, 2},
		{1, 4},
		{1, 5},
		{1, 8},
		{2, 8},
		{3, 5},
		{4, 1},
		{4, 7},
		{1, 6},
		{6, 2},
		{6, 4},
		{7, 4},
		{5, 8},
	}
	Q := []int64{0, 8, 4, 1, 12}
	expect := []int{4, 5, 5, 4, 6}
	runSample(t, n, E, Q, expect)
}
